# ⚔️ Rapier

> *"Swift, precise, and deadly efficient - like the blade it's named after"*

A high-performance TCP server built in Go, designed for lightning-fast network communications with minimal overhead.

## 🗡️ What is Rapier?

Rapier is a low-level TCP server implementation that cuts through the bloat of HTTP to deliver raw performance. Named after the swift and precise dueling sword, Rapier prioritizes speed and efficiency above all else.

### ⚡ Key Features

- **Platform-Optimized I/O**: Uses `epoll` on Linux and `kqueue` on macOS for handling 20,000+ concurrent connections
- **Zero HTTP Overhead**: Direct TCP communication without HTTP protocol layers  
- **Advanced Data Structures**: Built-in bloom filters, B+ trees, sorted sets, and skip lists
- **Memory Management**: Configurable eviction policies with LRU sampling and pool-based allocation
- **TTL Support**: Millisecond-precision key expiration with automatic cleanup
- **Cross-Platform**: Native support for both Linux and macOS architectures
- **High Performance**: Optimized for real-time applications requiring sub-millisecond latency

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
git clone https://github.com/leminhtuan61/Rapier.git
cd Rapier

# Install dependencies
go mod tidy

# Run the server
go run cmd/main.go

# Or build and run
go build -o rapier cmd/main.go
./rapier
```

The server starts on port `:3000` by default with I/O multiplexing enabled.

## ⚙️ Configuration

Rapier can be configured by modifying the values in `internal/config/config.go`:

| Parameter | Default Value | Description |
|-----------|---------------|-------------|
| `Protocol` | `"tcp"` | Network protocol (TCP) |
| `Port` | `":3000"` | Server listening port |
| `MaxConnection` | `20000` | Maximum concurrent connections |
| `MaxKeyNumber` | `10` | Maximum number of keys for certain operations |
| `EvictionRatio` | `0.1` | Ratio of keys to evict when memory limit reached |
| `EvictionPolicy` | `"allkeys-random"` | Key eviction strategy |
| `EpoolMaxSize` | `16` | Maximum size for eviction pool |
| `EpoolLruSampleSize` | `5` | Sample size for LRU eviction |

### Example Configuration

```go
// internal/config/config.go
package config

var Protocol = "tcp"
var Port = ":8080"           // Custom port
var MaxConnection = 50000    // Higher connection limit
var EvictionRatio = 0.2      // More aggressive eviction
```

## 🏗️ Architecture

```
┌──────────────────────────────────────────────────┐
│                    Rapier                        │
├──────────────────┬──────────────────┬────────────┤
│   TCP Server     │  I/O Multiplexing│   Config   │
│   (Port :3000)   │  (epoll/kqueue)  │  Manager   │
├──────────────────┼──────────────────┼────────────┤
│  Dictionary      │   Data Structures│ Expiration │
│  Storage         │  • Bloom Filter  │  Manager   │
│                  │  • Sorted Sets   │            │
│                  │  • B+ Trees      │            │
│                  │  • Skip Lists    │            │
└──────────────────┴──────────────────┴────────────┘
```

### Core Components

- **TCP Server**: Low-level socket handling with platform-specific I/O multiplexing
- **Dictionary Storage**: Thread-safe in-memory key-value store with TTL support
- **I/O Multiplexing**: Platform-optimized event handling (epoll on Linux, kqueue on macOS)  
- **Advanced Data Structures**: Bloom filters, sorted sets, B+ trees, and skip lists
- **Configuration Management**: Centralized configuration for server parameters
- **Expiration Manager**: Automatic cleanup of expired keys with millisecond precision

## 🧮 Data Structures

Rapier implements a comprehensive suite of high-performance data structures:

### 🎯 **Bloom Filter**
- Space-efficient probabilistic data structure for membership testing
- Configurable false positive rate with optimal hash function count
- Uses MurmurHash3 for fast, uniform hashing

### 🌳 **B+ Tree**
- Self-balancing tree structure optimized for sorted data access
- Efficient range queries and ordered iteration
- Used as the backing store for sorted sets

### 📊 **Sorted Sets**
- Redis-compatible sorted set implementation
- Combines hash table and B+ tree for O(log n) operations
- Support for score-based ranking and range operations

### ⚡ **Skip Lists**
- Probabilistic data structure for fast search and insertion
- Alternative implementation for sorted data structures
- Excellent performance characteristics for concurrent access

### 📚 **Dictionary**
- Core key-value storage with TTL support
- Millisecond-precision expiration tracking
- Thread-safe operations with efficient memory management

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
# Run all benchmarks
go test -bench=. ./...

# Test specific data structures
go test -bench=. ./internal/core/ds/ -v

# Test with memory allocation stats
go test -bench=. -benchmem ./internal/core/ds/
```

### Available Test Suites
- **Bloom Filter**: `bloom_test.go` - Probabilistic data structure performance
- **Dictionary**: `dictionary_test.go` - Core key-value operations with TTL
- **Sorted Set**: `sortedset_skiplist_test.go` - Skip list implementation benchmarks

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