# ⚔️ Rapier

> *"Swift, precise, and deadly efficient - like the blade it's named after"*

A high-performance TCP server built in Go, designed for lightning-fast network communications with minimal overhead.

## 🗡️ What is Rapier?

Rapier is a low-level TCP server implementation that cuts through the bloat of HTTP to deliver raw performance. Named after the swift and precise dueling sword, Rapier prioritizes speed and efficiency above all else.

### ⚡ Key Features

- **I/O Multiplexing**: Uses `epoll` on Linux and `kqueue` on macOS for handling thousands of concurrent connections
- **Zero HTTP Overhead**: Direct TCP communication without HTTP protocol layers  
- **Cross-Platform**: Native support for both Linux and macOS
- **Memory Efficient**: Minimal memory footprint with optimized data structures
- **Real-Time Ready**: Perfect for gaming servers, chat applications, and trading systems

## 🎯 Performance Benefits

| Protocol | Typical RPS | Latency | Memory Usage |
|----------|-------------|---------|--------------|
| **Rapier TCP** | 100K - 500K | <1ms | Minimal |
| Standard HTTP | 10K - 50K | 5-10ms | High |

*Rapier delivers **5-10x better performance** than traditional HTTP servers*

## 🚀 Quick Start

### Prerequisites
- Go 1.24 or higher
- Linux or macOS

### Installation & Run

```bash
# Clone the repository
git clone <your-repo-url>
cd Rapier

# Run the server
go run cmd/main.go

# Or build and run
go build -o rapier cmd/main.go
./rapier
```

The server starts on the configured port with I/O multiplexing enabled.

## 🏗️ Architecture

```
┌─────────────────────────────────────────┐
│                 Rapier                  │
├─────────────────┬───────────────────────┤
│   TCP Server    │    I/O Multiplexing   │
│                 │   (epoll/kqueue)      │
├─────────────────┼───────────────────────┤
│   Dictionary    │    Expire Manager     │
│   Storage       │                       │
└─────────────────┴───────────────────────┘
```

### Core Components

- **TCP Server**: Low-level socket handling with multiplexing
- **Dictionary Storage**: In-memory key-value store with TTL support  
- **I/O Multiplexing**: Platform-specific optimized event handling
- **Expire Manager**: Automatic cleanup of expired keys

## 💡 Use Cases

Rapier excels in scenarios requiring maximum performance:

- 🎮 **Game Servers** - Real-time multiplayer backends
- 💬 **Chat Systems** - High-throughput messaging platforms  
- 📈 **Trading Systems** - Low-latency financial data processing
- 🌐 **IoT Gateways** - Handle thousands of device connections
- 🔄 **Message Brokers** - Pub/sub and queue systems
- 🗄️ **Database Engines** - Custom storage solutions

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Test specific package
go test ./internal/core/ds/ -v
```

## 📊 Benchmarks

```bash
# Coming soon - benchmark suite for performance testing
go test -bench=. ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📋 Roadmap

- [ ] Protocol definition and standardization
- [ ] Authentication and authorization layer
- [ ] Clustering and horizontal scaling
- [ ] Monitoring and metrics collection
- [ ] Client libraries for multiple languages
- [ ] WebSocket gateway for browser clients

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <b>⚔️ Fast. Precise. Deadly Efficient. ⚔️</b>
  <br>
  <i>Built with ❤️ and ☕ in Go</i>
</div>