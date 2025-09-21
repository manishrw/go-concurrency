package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// === Basic Goroutines ===
	// This module demonstrates fundamental goroutine concepts in Go.
	// Complete the following exercises:

	// TODO: Implement the following basic goroutine patterns:

	// 1. Basic Goroutine Creation
	//    - Create simple goroutines
	//    - Understand goroutine lifecycle
	//    - Handle goroutine execution order
	//    - Use time.Sleep for synchronization

	// 2. Goroutine with Parameters
	//    - Pass parameters to goroutines
	//    - Handle closure variable capture
	//    - Avoid common goroutine pitfalls
	//    - Use proper variable scoping

	// 3. Multiple Goroutines
	//    - Create multiple concurrent goroutines
	//    - Understand goroutine scheduling
	//    - Handle non-deterministic execution
	//    - Use sync.WaitGroup for coordination

	// 4. Goroutine Communication
	//    - Use channels for goroutine communication
	//    - Implement basic producer-consumer
	//    - Handle channel blocking
	//    - Understand channel direction

	// 5. Goroutine Error Handling
	//    - Handle errors in goroutines
	//    - Implement panic recovery
	//    - Use error channels
	//    - Handle goroutine failures

	// 6. Goroutine Best Practices
	//    - Avoid goroutine leaks
	//    - Implement proper cleanup
	//    - Use context for cancellation
	//    - Handle resource management

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining concepts
	// - Test with different scenarios

	fmt.Println("=== Basic Goroutines ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each goroutine pattern!")
	fmt.Println()

	// Example implementations (to be replaced with your code):
	fmt.Println("Example implementations:")

	// Example 1: Basic goroutine
	fmt.Println("\n1. Basic Goroutine Example:")
	go func() {
		fmt.Println("Hello from a goroutine!")
	}()
	time.Sleep(100 * time.Millisecond)

	// Example 2: Multiple goroutines with WaitGroup
	fmt.Println("\n2. Multiple Goroutines Example:")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is running\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed!")
}
