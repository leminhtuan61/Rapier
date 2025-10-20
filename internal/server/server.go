package server

import (
	"Nietzsche/internal/config"
	"Nietzsche/internal/constant"
	"Nietzsche/internal/core"
	"Nietzsche/internal/core/io_multiplexing"
	"hash/fnv"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var serverStatus int32 = constant.ServerStatusIdle

func readCommand(fd int) (*core.Command, error) {
	var buf = make([]byte, 512)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, io.EOF
	}
	return core.ParseCmd(buf)
}

func readCommandConn(conn net.Conn) (*core.Command, error) {
	var buf = make([]byte, 512)
	// Use the Read method from the net.Conn interface
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err // This will properly handle io.EOF
	}
	return core.ParseCmd(buf[:n])
}

func respond(data string, fd int) error {
	if _, err := syscall.Write(fd, []byte(data)); err != nil {
		return err
	}
	return nil
}

func WaitForSignal(wg *sync.WaitGroup, signals chan os.Signal) {
	defer wg.Done()
	// Wait for signal in channel, it not available then wait
	<-signals
	// Busy loop
	for {
		if atomic.CompareAndSwapInt32(&serverStatus, constant.ServerStatusIdle, constant.ServerStatusShuttingDown) {
			// The swap was successful! We have now claimed the shutdown state.
			log.Println("Shutting down gracefully")
			os.Exit(0)
		}
	}
}

type Server struct {
	workers       []*core.Worker
	ioHandlers    []*IOHandler
	numWorkers    int
	numIOHandlers int

	// For round-robin assigment of new connection to I/O handlers
	nextIOHandler int
}

func (s *Server) getPartitionID(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % s.numWorkers
}

// set abc 123
// abc -> 1
// get abc
// abc -> 1
func (s *Server) dispatch(task *core.Task) {
	// Commands like PING etc., don't have a key.
	// We can send them to any worker.
	var key string
	var workerID int
	if len(task.Command.Args) > 0 {
		key = task.Command.Args[0]
		workerID = s.getPartitionID(key)
	} else {
		workerID = rand.Intn(s.numWorkers)
	}

	s.workers[workerID].TaskCh <- task
}

func NewServer() *Server {
	numCores := runtime.NumCPU()  // 8
	numIOHandlers := numCores / 2 // 4
	numWorkers := numCores / 2    // 4
	log.Printf("Initializing server with %d workers and %d io handler\n", numWorkers, numIOHandlers)

	s := &Server{
		workers:       make([]*core.Worker, numWorkers),
		ioHandlers:    make([]*IOHandler, numIOHandlers),
		numWorkers:    numWorkers,
		numIOHandlers: numIOHandlers,
	}

	for i := 0; i < numWorkers; i++ {
		s.workers[i] = core.NewWorker(i, 1024)
	}

	for i := 0; i < numIOHandlers; i++ {
		handler, err := NewIOHandler(i, s)
		if err != nil {
			log.Fatalf("Failed to create I/O handler %d: %v", i, err)
		}
		s.ioHandlers[i] = handler
	}
	return s
}

func RunIoMultiplexingServer(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("starting an I/O Multiplexing TCP server on", config.Port)
	listener, err := net.Listen(config.Protocol, config.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Get the file descriptor from the listener
	tcpListener, ok := listener.(*net.TCPListener)
	if !ok {
		log.Fatal("listener is not a TCPListener")
	}
	listenerFile, err := tcpListener.File()
	if err != nil {
		log.Fatal(err)
	}
	defer listenerFile.Close()

	serverFd := int(listenerFile.Fd())

	// Create an ioMultiplexer instance (epoll in Linux, kqueue in MacOS)
	ioMultiplexer, err := io_multiplexing.CreateIOMultiplexer()
	if err != nil {
		log.Fatal(err)
	}
	defer ioMultiplexer.Close()

	// Monitor "read" events on the Server FD
	if err = ioMultiplexer.Monitor(io_multiplexing.Event{
		Fd: serverFd,
		Op: io_multiplexing.OpRead,
	}); err != nil {
		log.Fatal(err)
	}

	var events = make([]io_multiplexing.Event, config.MaxConnection)
	var lastActiveExpireExecTime = time.Now()
	for atomic.LoadInt32(&serverStatus) != constant.ServerStatusShuttingDown {
		// Check last execution time and call if it is more than 100ms ago.
		if time.Now().After(lastActiveExpireExecTime.Add(constant.ActiveExpireFrequency)) {
			if !atomic.CompareAndSwapInt32(&serverStatus, constant.ServerStatusIdle, constant.ServerStatusBusy) {
				if serverStatus == constant.ServerStatusShuttingDown {
					return
				}
			}
			core.ActiveDeleteExpiredKeys() // Busy
			atomic.SwapInt32(&serverStatus, constant.ServerStatusIdle)
			// Idle
			lastActiveExpireExecTime = time.Now()
		}
		// wait for file descriptors in the monitoring list to be ready for I/O
		// it is a blocking call.
		// Idle
		events, err = ioMultiplexer.Wait()
		if err != nil {
			continue
		}
		// Goroutine #2 is gracefully shutdown
		// means: serverStatus == ServerStatusShuttingDown
		if !atomic.CompareAndSwapInt32(&serverStatus, constant.ServerStatusIdle, constant.ServerStatusBusy) {
			if serverStatus == constant.ServerStatusShuttingDown {
				return
			}
		}
		// Busy
		for i := 0; i < len(events); i++ {
			if events[i].Fd == serverFd {
				log.Printf("new client is trying to connect")
				// set up new connection
				connFd, _, err := syscall.Accept(serverFd)
				if err != nil {
					log.Println("err", err)
					continue
				}
				log.Printf("set up a new connection")
				// ask epoll to monitor this connection
				if err = ioMultiplexer.Monitor(io_multiplexing.Event{
					Fd: connFd,
					Op: io_multiplexing.OpRead,
				}); err != nil {
					log.Fatal(err)
				}
			} else {
				cmd, err := readCommand(events[i].Fd)
				if err != nil {
					if err == io.EOF || err == syscall.ECONNRESET {
						log.Println("client disconnected")
						_ = syscall.Close(events[i].Fd)
						continue
					}
					log.Println("read error:", err)
					continue
				}
				if err = core.ExecuteAndResponse(cmd, events[i].Fd); err != nil {
					log.Println("err write:", err)
				}
			}
		}
		// Idle
		atomic.SwapInt32(&serverStatus, constant.ServerStatusIdle)
	}
}
