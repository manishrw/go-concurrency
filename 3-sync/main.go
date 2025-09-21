package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// === Synchronization Primitives ===
	// This module demonstrates synchronization primitives in Go.
	// Complete the following exercises:

	// TODO: Implement the following synchronization patterns:

	// 1. Mutex and RWMutex
	//    - Use sync.Mutex for exclusive access
	//    - Use sync.RWMutex for read-write operations
	//    - Understand lock contention
	//    - Implement proper lock ordering

	// 2. WaitGroup
	//    - Use sync.WaitGroup for goroutine coordination
	//    - Handle dynamic goroutine counts
	//    - Implement proper Add/Done/Wait usage
	//    - Avoid common WaitGroup pitfalls

	// 3. Once
	//    - Use sync.Once for one-time initialization
	//    - Implement singleton patterns
	//    - Handle initialization errors
	//    - Understand Once behavior

	// 4. Cond (Condition Variables)
	//    - Use sync.Cond for condition-based waiting
	//    - Implement producer-consumer with conditions
	//    - Handle spurious wakeups
	//    - Use Signal() and Broadcast()

	// 5. Map
	//    - Use sync.Map for concurrent map operations
	//    - Compare with regular map + mutex
	//    - Implement concurrent cache
	//    - Handle map operations safely

	// 6. Pool
	//    - Use sync.Pool for object reuse
	//    - Implement connection pooling
	//    - Handle pool exhaustion
	//    - Implement custom pool cleanup

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining concepts
	// - Test with different scenarios

	fmt.Println("=== Synchronization Primitives ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each synchronization pattern!")
	fmt.Println()

	// Example implementations (to be replaced with your code):
	fmt.Println("Example implementations:")

	// Example 1: Mutex
	fmt.Println("\n1. Mutex Example:")
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)

	// Example 2: RWMutex
	fmt.Println("\n2. RWMutex Example:")
	var (
		data    = make(map[string]int)
		rwMutex sync.RWMutex
	)

	// Writer goroutine
	go func() {
		for i := 0; i < 5; i++ {
			rwMutex.Lock()
			data[fmt.Sprintf("key%d", i)] = i * 10
			rwMutex.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Reader goroutines
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				rwMutex.RLock()
				fmt.Printf("Reader %d: data = %v\n", id, data)
				rwMutex.RUnlock()
				time.Sleep(30 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)

	// Example 3: Once
	fmt.Println("\n3. sync.Once Example:")
	var once sync.Once
	initFunc := func() {
		fmt.Println("This will only be called once!")
	}

	for i := 0; i < 5; i++ {
		go once.Do(initFunc)
	}

	time.Sleep(100 * time.Millisecond)
}
