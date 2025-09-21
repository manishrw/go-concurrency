package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// === Worker Pool Patterns ===
	// This module demonstrates various worker pool implementations in Go.
	// Complete the following exercises:

	// TODO: Implement the following worker pool patterns:

	// 1. Basic Worker Pool
	//    - Create fixed number of workers
	//    - Use channels for job distribution
	//    - Implement result collection
	//    - Handle worker lifecycle

	// 2. Dynamic Worker Pool
	//    - Scale workers based on load
	//    - Implement worker health checks
	//    - Handle worker failures
	//    - Implement graceful scaling

	// 3. Priority Worker Pool
	//    - Implement priority queues
	//    - Handle high-priority jobs first
	//    - Implement job preemption
	//    - Handle priority inversion

	// 4. Worker Pool with Backpressure
	//    - Implement flow control
	//    - Handle queue overflow
	//    - Implement rate limiting
	//    - Handle backpressure scenarios

	// 5. Worker Pool with Rate Limiting
	//    - Implement rate limiting
	//    - Handle burst traffic
	//    - Implement token bucket
	//    - Handle rate limit violations

	// 6. Worker Pool Best Practices
	//    - Implement proper error handling
	//    - Add monitoring and metrics
	//    - Handle graceful shutdown
	//    - Implement resource cleanup

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining concepts
	// - Test with different scenarios

	fmt.Println("=== Worker Pool Patterns ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each worker pool pattern!")
	fmt.Println()

	// Example implementations (to be replaced with your code):
	fmt.Println("Example implementations:")

	// Example 1: Basic worker pool
	fmt.Println("\n1. Basic Worker Pool:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	numWorkers := 3
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(500 * time.Millisecond) // Simulate work
				results <- job * 2
			}
		}(i)
	}

	// Send jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	// Example 2: Worker pool with rate limiting
	fmt.Println("\n2. Worker Pool with Rate Limiting:")
	jobs2 := make(chan string, 10)
	rateLimiter := time.Tick(200 * time.Millisecond) // Process one job every 200ms

	go func() {
		for i := 1; i <= 5; i++ {
			jobs2 <- fmt.Sprintf("task-%d", i)
		}
		close(jobs2)
	}()

	for job := range jobs2 {
		<-rateLimiter // Wait for rate limit
		fmt.Printf("Processing %s\n", job)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("All worker pool examples completed!")
}
