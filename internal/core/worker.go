package core

import (
	"Nietzsche/internal/constant"
	"Nietzsche/internal/data_structure"
	"errors"
	"fmt"
	"strconv"
)

type Task struct {
	Command *Command
	ReplyCh chan []byte // Channel to send the result back to the client's handler
}

type Worker struct {
	id        int
	dictStore *data_structure.Dict
	TaskCh    chan *Task // Receives tasks from the I/O handler
}

func NewWorker(id int, bufferSize int) *Worker {
	w := &Worker{
		id:        id,
		dictStore: data_structure.CreateDict(),
		TaskCh:    make(chan *Task, bufferSize),
	}
	go w.run() // new routine
	return w
}

func (w *Worker) cmdSET(args []string) []byte {
	if len(args) < 2 || len(args) == 3 || len(args) > 4 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'SET' command"), false)
	}

	var key, value string
	var ttlMs int64 = -1

	key, value = args[0], args[1]
	if len(args) > 2 {
		ttlSec, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return Encode(errors.New("(error) ERR value is not an integer or out of range"), false)
		}
		ttlMs = ttlSec * 1000
	}

	w.dictStore.Set(key, w.dictStore.NewObj(key, value, ttlMs))
	return constant.RespOk
}

func (w *Worker) cmdPING(args []string) []byte {
	var res []byte
	if len(args) > 1 {
		return Encode(errors.New("ERR wrong number of arguments for 'ping' command"), false)
	}

	if len(args) == 0 {
		res = Encode("PONG", true)
	} else {
		res = Encode(args[0], false)
	}
	return res
}

func (w *Worker) cmdGET(args []string) []byte {
	if len(args) != 1 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'GET' command"), false)
	}
	key := args[0]
	obj := w.dictStore.Get(key)
	if obj == nil {
		return constant.RespNil
	}

	if w.dictStore.HasExpired(key) {
		return constant.RespNil
	}

	return Encode(obj.Value, false)
}

func (w *Worker) ExecuteAndResponse(task *Task) {
	//log.Printf("worker %d executes command %s", w.id, task.Command)
	var res []byte

	switch task.Command.Cmd {
	case "SET":
		res = w.cmdSET(task.Command.Args)
	case "GET":
		res = w.cmdGET(task.Command.Args)
	case "PING":
		res = w.cmdPING(task.Command.Args)
	default:
		res = []byte(fmt.Sprintf("-CMD NOT FOUND\r\n"))
	}
	task.ReplyCh <- res
}

func (w *Worker) run() {
	for task := range w.TaskCh {
		w.ExecuteAndResponse(task)
	}
}
