# ⚔️ Rapier

> *"Swift, precise, and deadly efficient - like the blade it's named after"*

A high-performance Redis-like server built from scratch in Go, designed to handle millions of requests with ultra-low latency.

## 🗡️ What is Rapier?

Rapier is a Redis server implementation built from the ground up, optimized for maximum performance. Named after the swift and precise dueling sword, Rapier prioritizes speed and efficiency above all else.

### ⚡ Key Features

- **Platform-Optimized I/O**: Uses `epoll` on Linux and `kqueue` on macOS for handling 20,000+ concurrent connections
- **Event-driven Architecture**: Similar to Node.js but with higher performance
- **Advanced Data Structures**: Built-in bloom filters, B+ trees, sorted sets, and skip lists
- **Memory Management**: Configurable eviction policies with LRU sampling and pool-based allocation
- **TTL Support**: Millisecond-precision key expiration with automatic cleanup
- **Cross-Platform**: Native support for both Linux and macOS architectures
- **High Performance**: Optimized for real-time applications requiring sub-millisecond latency

## 🎯 Performance Benefits

| Protocol | Typical RPS | Latency | Memory Usage |
|----------|-------------|---------|--------------|
| **Rapier Redis** | 120K+ | <0.3ms | Minimal |
| Standard Redis | 50K-80K | 1-2ms | Medium |
| HTTP Server | 10K-20K | 5-10ms | High |

*Rapier delivers **2-3x better performance** than standard Redis*

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- Linux or macOS
- Redis client tools (for testing)

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

The server starts on port `:9000` by default with I/O multiplexing enabled.

## ⚙️ Configuration

Rapier can be configured by modifying the values in `internal/config/config.go`:

| Parameter | Default Value | Description |
|-----------|---------------|-------------|
| `Port` | `":9000"` | Server listening port |
| `MaxConnection` | `20000` | Maximum concurrent connections |
| `MaxKeyNumber` | `1000000` | Maximum number of keys for certain operations |
| `EvictionRatio` | `0.1` | Ratio of keys to evict when memory limit reached |
| `EvictionPolicy` | `"allkeys-lru"` | Key eviction strategy |
| `ListenerNumber` | `2` | Number of listeners |

### Example Configuration

```go
// internal/config/config.go
var Port = ":8080"           // Custom port
var MaxConnection = 50000    // Higher connection limit
var EvictionRatio = 0.2      // More aggressive eviction
```

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                        Rapier                            │
├─────────────────┬─────────────────┬─────────────────────┤
│   TCP Server    │ I/O Multiplexing│     Data Store      │
│   (Port :9000)  │ (epoll/kqueue)  │  • Dictionary       │
├─────────────────┼─────────────────┼  • Sorted Sets      │
│  Event Loop     │   Workers       │  • Bloom Filters    │
│  (Node.js-like) │  (Goroutines)   │  • B+ Trees         │
├─────────────────┼─────────────────┼─────────────────────┤
│  Redis Protocol │   TTL Manager   │   Advanced DS       │
│  (RESP)         │   (Expiration)  │  • Skip Lists       │
└─────────────────┴─────────────────┴─────────────────────┘
```

### Core Components

- **TCP Server**: Low-level socket handling with platform-specific I/O multiplexing
- **Event Loop**: Event-driven architecture similar to Node.js
- **I/O Multiplexing**: Platform-optimized event handling (epoll on Linux, kqueue on macOS)
- **Data Structures**: Bloom filters, sorted sets, B+ trees, and skip lists
- **TTL Manager**: Automatic cleanup of expired keys with millisecond precision

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

## 🎮 Supported Redis Commands

### Basic Commands
- `PING` - Test connection
- `SET key value` - Store key-value
- `GET key` - Retrieve value
- `TTL key` - Check time to live

### Set Operations
- `SADD key member1 member2...` - Add members to set
- `SREM key member1 member2...` - Remove members from set
- `SMEMBERS key` - Get all members
- `SISMEMBER key member` - Check if member exists in set

### Sorted Set Operations
- `ZADD key score1 member1 score2 member2...` - Add to sorted set
- `ZSCORE key member` - Get member score
- `ZRANK key member` - Get member rank

### Bloom Filter Operations
- `BF.RESERVE key capacity error_rate` - Create bloom filter
- `BF.MADD key item1 item2...` - Add items to bloom filter
- `BF.EXISTS key item` - Check if item exists in bloom filter

### Count-Min Sketch Operations
- `CMS.INITBYDIM key width depth` - Initialize CMS
- `CMS.INITBYPROB key error_rate probability` - Initialize CMS by probability
- `CMS.INCRBY key item count` - Increment item count
- `CMS.QUERY key item` - Query item count

## 🧪 Testing

### Run Test Suite

```bash
# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Test specific package
go test ./internal/data_structure/ -v
```

### Benchmark Tests

```bash
# Run all benchmarks
go test -bench=. ./...

# Test specific data structures
go test -bench=. ./internal/data_structure/ -v

# Test with memory allocation stats
go test -bench=. -benchmem ./internal/data_structure/
```

## 📊 Performance Benchmarks

### Real Benchmark Results

```bash
# Test SET operations
redis-benchmark -p 9000 -t set -n 100000 -r 100000

# Test GET operations
redis-benchmark -p 9000 -t get -n 100000 -r 100000

# Test with multi-threading
redis-benchmark -p 9000 -t set,get -n 100000 --threads 3
```

**Results:**
- **Throughput**: 120,627 requests/second
- **Average Latency**: 0.253ms
- **P50 Latency**: 0.247ms
- **P95 Latency**: 0.391ms
- **P99 Latency**: 0.775ms

### Performance Comparison

| Test Type | Throughput | Avg Latency | P50 | P95 | P99 |
|-----------|------------|-------------|-----|-----|-----|
| **SET (Single)** | 120,627 RPS | 0.253ms | 0.247ms | 0.391ms | 0.775ms |
| **GET (Single)** | 116,959 RPS | 0.265ms | 0.247ms | 0.663ms | 0.807ms |
| **SET (Multi-thread)** | 100,000 RPS | 0.315ms | 0.263ms | 0.703ms | 0.895ms |
| **GET (Multi-thread)** | 133,333 RPS | 0.307ms | 0.271ms | 0.591ms | 0.823ms |

## 💡 Use Cases

Rapier excels in scenarios requiring maximum performance:

- 🎮 **Game Servers** - Real-time multiplayer backends
- 💬 **Chat Systems** - High-throughput messaging platforms
- 📈 **Trading Systems** - Low-latency financial data processing
- 🌐 **IoT Gateways** - Handle thousands of device connections
- 🔄 **Message Brokers** - Pub/sub and queue systems
- 🗄️ **Database Engines** - Custom storage solutions

## 🔧 Usage

### Connect with Redis Client

```bash
# Using redis-cli
redis-cli -p 9000

# Test connection
PING

# Test basic commands
SET mykey "Hello Rapier"
GET mykey

# Test Set operations
SADD myset "apple" "banana" "cherry"
SMEMBERS myset

# Test Sorted Set operations
ZADD leaderboard 100 "player1" 200 "player2"
ZSCORE leaderboard "player1"
```

### Monitoring and Profiling

```bash
# Profiling endpoint
curl http://localhost:6060/debug/pprof/

# Monitor system resources
htop

# Check server status
lsof -i :9000
```

## 🚀 I/O Multiplexing Architecture

### Event Loop Pattern (Node.js-like)

```go
// Main event loop
for serverStatus != ShuttingDown {
    // 1. Wait for events (blocking call)
    events, err = ioMultiplexer.Wait()
    
    // 2. Process events
    for _, event := range events {
        if event.Fd == serverFd {
            // Accept new connection
        } else {
            // Handle client data
        }
    }
}
```

### Cross-Platform Support

- **Linux**: Uses `epoll` (similar to Node.js on Linux)
- **macOS**: Uses `kqueue` (similar to Node.js on macOS)
- **Interface abstraction**: Common interface for both platforms

### Multi-threading Architecture

```go
// Multiple I/O handlers
for _, handler := range ioHandlers {
    go handler.Run()  // Each handler has its own event loop
}

// Round-robin load balancing
handler := ioHandlers[nextHandler % numHandlers]
```

## 🎯 I/O Multiplexing vs Node.js

| **Node.js** | **Rapier (Go)** |
|-------------|-----------------|
| `libuv` event loop | `epoll`/`kqueue` |
| `setImmediate()` queue | Goroutine scheduling |
| `process.nextTick()` | Channel communication |
| Callback execution | Event processing loop |
| Single-threaded | I/O multiplexing + workers |

## 📋 Roadmap

- [ ] Protocol definition and standardization
- [ ] Authentication and authorization layer
- [ ] Clustering and horizontal scaling
- [ ] Monitoring and metrics collection
- [ ] Client libraries for multiple languages
- [ ] WebSocket gateway for browser clients
- [ ] Replication and persistence
- [ ] Advanced data structures (HyperLogLog, Bitmaps)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <b>⚔️ Fast. Precise. Deadly Efficient. ⚔️</b>
  <br>
  <i>Built with ❤️ and ☕ in Go</i>
</div>