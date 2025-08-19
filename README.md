# Rapier - TCP Server Implementations in Go

A comprehensive Go project demonstrating three different approaches to building high-performance TCP servers. This project showcases various concurrency models and network programming techniques, from basic goroutines to advanced I/O multiplexing.

## 🏗️ Project Structure

```
Rapier/
├── TcpServer/           # Basic goroutine-per-connection server
├── ThreadPool/          # Worker pool pattern server  
├── cmd/                 # I/O multiplexing server (main implementation)
├── internal/            # Core libraries and utilities
│   ├── config/          # Server configuration
│   ├── server/          # Main server implementation
│   └── core/            # RESP protocol & I/O multiplexing
└── doc/                 # Documentation
```

## 🚀 Server Implementations

### 1. Basic TCP Server (`TcpServer/`)
**Approach**: One goroutine per connection
- Simple HTTP-like responses
- 10-second processing delay simulation
- Unlimited goroutines (can lead to resource exhaustion)

```bash
go run TcpServer/main.go
```

### 2. Thread Pool Server (`ThreadPool/`)
**Approach**: Fixed worker pool pattern
- 2 worker goroutines handling job queue
- Controlled resource usage
- HTTP-like responses with 1-second delay

```bash
go run ThreadPool/main.go
```

### 3. I/O Multiplexing Server (`cmd/`)
**Approach**: Event-driven I/O multiplexing
- **Linux**: epoll-based
- **macOS**: kqueue-based  
- Supports up to 20,000 concurrent connections
- Echo server functionality
- RESP protocol support

```bash
go run cmd/main.go
```

## 🔧 Features by Implementation

| Feature | Basic TCP | Thread Pool | I/O Multiplexing |
|---------|-----------|-------------|------------------|
| Max Connections | Unlimited* | Limited by pool | 20,000 |
| Resource Usage | High | Controlled | Very Low |
| Scalability | Poor | Good | Excellent |
| Complexity | Simple | Medium | Advanced |
| Protocol | HTTP-like | HTTP-like | Echo/RESP |

*Limited by system resources

## 🛠️ Technical Components

### I/O Multiplexing (`internal/core/io_multiplexing/`)
Platform-specific implementations:
- **epoll** (Linux): Edge-triggered notifications
- **kqueue** (macOS): BSD-style event notifications
- Cross-platform interface for portability

### RESP Protocol (`internal/core/resp.go`)
Redis Serialization Protocol implementation:
- Simple Strings: `+OK\r\n`
- Integers: `:123\r\n`  
- Bulk Strings: `$5\r\nhello\r\n`
- Arrays: `*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n`
- Errors: `-Error message\r\n`

### Configuration (`internal/config/`)
- Protocol: TCP
- Port: 3000
- Max Connections: 20,000

## 📊 Performance Comparison

### Basic TCP Server
- **Pros**: Simple to understand and implement
- **Cons**: Memory usage grows with connections, potential for goroutine exhaustion
- **Use Case**: Low-traffic applications, prototypes

### Thread Pool Server  
- **Pros**: Controlled resource usage, good for CPU-bound tasks
- **Cons**: Limited by pool size, potential queue buildup
- **Use Case**: Moderate traffic with predictable load

### I/O Multiplexing Server
- **Pros**: Handles thousands of connections efficiently, minimal memory per connection
- **Cons**: Complex implementation, harder to debug
- **Use Case**: High-traffic applications, real-time systems

## 🚦 Getting Started

### Prerequisites
- Go 1.24+
- Linux or macOS (for I/O multiplexing)

### Quick Start

1. **Clone and setup**:
```bash
git clone <repository-url>
cd Rapier
go mod tidy
```

2. **Run any server**:
```bash
# Basic server
go run TcpServer/main.go

# Thread pool server
go run ThreadPool/main.go

# I/O multiplexing server (recommended)
go run cmd/main.go
```

3. **Test the server**:
```bash
# Multiple connections
for i in {1..10}; do echo "Hello $i" | nc localhost 3000 & done

# Single connection
echo "Hello World" | nc localhost 3000
```

## 🧪 Testing & Benchmarking

```bash
# Run tests
go test ./...

# Benchmark with multiple connections
# Terminal 1: Start server
go run cmd/main.go

# Terminal 2: Load test
for i in {1..1000}; do echo "test $i" | nc localhost 3000 & done
```

## 📚 Learning Objectives

This project demonstrates:
- **Concurrency Patterns**: Goroutines, channels, worker pools
- **Network Programming**: TCP sockets, system calls
- **I/O Models**: Blocking, non-blocking, event-driven
- **Protocol Implementation**: HTTP-like responses, RESP protocol
- **System Programming**: epoll, kqueue, file descriptors
- **Performance Optimization**: Memory efficiency, connection scaling

## 🔍 Code Examples

### Basic Connection Handling
```go
// TcpServer approach
go handleConnection(conn)

// ThreadPool approach  
pool.AddJob(conn)

// I/O Multiplexing approach
ioMultiplexer.Monitor(Event{Fd: connFd, Op: OpRead})
```

### RESP Encoding
```go
// Encode string
resp := Encode("hello", false)  // $5\r\nhello\r\n

// Encode array
resp := Encode([]string{"hello", "world"}, false)
```

## 🤝 Contributing

This is an educational project demonstrating different server architectures. Feel free to:
- Add new server implementations
- Improve existing protocols
- Add benchmarking tools
- Enhance documentation

## 📄 License

Educational/demonstration purposes. Free to use and modify.

---

**Note**: This project is designed for learning systems programming concepts. The I/O multiplexing implementation showcases production-ready patterns used in high-performance servers like Redis and Nginx.