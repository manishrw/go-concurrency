# Go Concurrency Learning Path: SDE1 to Principal Engineer

This project provides a comprehensive learning path for Go concurrency, designed to take you from basic understanding to expert-level knowledge expected of a Principal Engineer.

## Learning Path Overview

### **Level 1: Foundation **
*Prerequisites: Basic Go knowledge*

- **1-`basic-goroutine/`** - Fundamental goroutine concepts
- **2-`channels/`** - Basic channel operations and patterns

### **Level 2: Intermediate **
*Prerequisites: Level 1 completion*

- **3-`sync/`** - Synchronization primitives (Mutex, RWMutex, WaitGroup, Once)
- **4-`context/`** - Context package for cancellation and timeouts
- **5-`channel-patterns/`** - Advanced channel patterns (select, fan-in/out, pipelines)
- **6-`error-handling/`** - Error handling in concurrent code
- **7-`sync-advanced/`** - Advanced sync primitives (Cond, Map, Pool)

### **Level 3: Advanced **
*Prerequisites: Level 2 completion*

- **8-`worker-pools/`** - Worker pool patterns and implementations
- **9-`pipeline-patterns/`** - Pipeline and streaming patterns
- **10-`performance-optimization/`** - Performance tuning and optimization
- **11-`testing-concurrency/`** - Testing strategies for concurrent code

### **Level 4: Expert **
*Prerequisites: Level 3 completion*

- **12-`lock-free-programming/`** - Atomic operations and lock-free algorithms
- **13-`complex-patterns/`** - Advanced concurrency patterns (Actor, CSP, Reactive)
- **14-`system-design/`** - Concurrency in system design and microservices
- **15-`advanced-topics/`** - Goroutine internals, memory model, debugging

## Quick Reference

### **All Modules (1-15)**
```
1-basic-goroutine/          - Fundamental goroutine concepts
2-channels/                 - Basic channel operations and patterns
3-sync/                     - Synchronization primitives
4-context/                  - Context package for cancellation and timeouts
5-channel-patterns/         - Advanced channel patterns
6-error-handling/           - Error handling in concurrent code
7-sync-advanced/            - Advanced sync primitives
8-worker-pools/             - Worker pool patterns and implementations
9-pipeline-patterns/        - Pipeline and streaming patterns
10-performance-optimization/ - Performance tuning and optimization
11-testing-concurrency/     - Testing strategies for concurrent code
12-lock-free-programming/   - Atomic operations and lock-free algorithms
13-complex-patterns/        - Advanced concurrency patterns
14-system-design/           - Concurrency in system design and microservices
15-advanced-topics/         - Goroutine internals, memory model, debugging
```

### **Quick Start Commands**
```bash
# Start with module 1
cd 1-basic-goroutine && go run main.go

# Jump to any module
cd 5-channel-patterns && go run main.go

# List all modules
ls -1 | grep "^[0-9]"
```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-concurrency
   ```

2. **Start with Level 1**
   ```bash
   cd 1-basic-goroutine
   go run main.go
   ```

3. **Follow the learning path**
   - Each module contains detailed instructions
   - Implement the exercises one by one
   - Test your understanding with different scenarios

## Learning Progression

### **Week 1-2 : Foundation**
- Complete **1-`basic-goroutine/`** and **2-`channels/`**
- Understand goroutine lifecycle and basic channel operations
- Practice with simple concurrent programs

### **Week 3-4 : Intermediate**
- Complete **3-`sync/`** and **4-`context/`**
- Learn synchronization primitives and context usage
- Implement proper error handling

### **Week 5-6 : Advanced Patterns**
- Complete **5-`channel-patterns/`** and **6-`error-handling/`**
- Master advanced channel patterns
- Implement robust error handling

### **Week 7-8 : Advanced Sync**
- Complete **7-`sync-advanced/`** and **8-`worker-pools/`**
- Learn advanced synchronization techniques
- Implement worker pool patterns

### **Week 9-10 (Senior): Performance & Testing**
- Complete **9-`pipeline-patterns/`** and **10-`performance-optimization/`**
- Learn performance optimization techniques
- Implement comprehensive testing

### **Week 11-12 : Advanced Topics**
- Complete **11-`testing-concurrency/`** and **12-`lock-free-programming/`**
- Master testing concurrent code
- Learn lock-free programming

### **Week 13-14 : Expert Level**
- Complete **13-`complex-patterns/`**, **14-`system-design/`**, and **15-`advanced-topics/`**
- Master complex concurrency patterns
- Understand system design with concurrency
- Learn Go internals and debugging

## Module Structure

Each module follows this structure:
- **`main.go`** - Contains exercises and examples
- **Instructions** - Detailed learning objectives
- **Examples** - Working code examples
- **Exercises** - Hands-on implementation tasks

## Prerequisites

- Go 1.19 or later
- Basic understanding of Go syntax
- Familiarity with basic programming concepts

## Useful Commands

```bash
# Run race detection
go test -race -v

# Run benchmarks
go test -bench=. -benchmem

# Profile CPU usage
go tool pprof

# Trace execution
go tool trace

# Check for goroutine leaks
go test -race -v -timeout=30s
```

## Best Practices

1. **Always use race detection** during development
2. **Implement proper error handling** in concurrent code
3. **Use context for cancellation** and timeouts
4. **Avoid goroutine leaks** with proper cleanup
5. **Test concurrent code thoroughly** with different scenarios
6. **Profile and benchmark** your concurrent code
7. **Document concurrency assumptions** and invariants

## Contributing

This learning path is designed to be hands-on. Each module contains:
- Clear learning objectives
- Working examples
- Hands-on exercises
- Best practices and common pitfalls

Feel free to extend the modules with additional examples and exercises!

## Resources

- [Go Memory Model](https://golang.org/ref/mem)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go.html#concurrency)
- [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide)
- [Advanced Go Concurrency Patterns](https://talks.golang.org/2013/advconc.slide)

---

**Happy Learning!** 🚀

Start with Level 1 and work your way up. Each module builds upon the previous ones, ensuring a solid foundation before moving to more complex topics.
