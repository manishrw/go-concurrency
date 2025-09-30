package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Basic Goroutines ===")
	fmt.Println("Demonstrating fundamental goroutine concepts in Go")
	fmt.Println()

	// 1. Basic Goroutine Creation
	basicGoroutineCreation()

	// 2. Goroutine with Parameters
	goroutineWithParameters()

	// 3. Multiple Goroutines
	multipleGoroutines()

	// 4. Goroutine Communication
	goroutineCommunication()

	// 5. Goroutine Error Handling
	goroutineErrorHandling()

	// 6. Goroutine Best Practices
	goroutineBestPractices()
}

// 1. Basic Goroutine Creation
// Demonstrates simple goroutine creation, lifecycle, and execution order
func basicGoroutineCreation() {
	fmt.Println("=== 1. Basic Goroutine Creation ===")

	// Simple goroutine creation
	fmt.Println("Creating a basic goroutine...")
	go func() {
		fmt.Println("  Hello from goroutine!")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("  Goroutine is finishing...")
	}()

	// Wait for goroutine to complete
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Basic goroutine completed!")
	fmt.Println()
}

// 2. Goroutine with Parameters
// Demonstrates passing parameters and avoiding closure variable capture issues
func goroutineWithParameters() {
	fmt.Println("=== 2. Goroutine with Parameters ===")

	// Correct way: Pass parameters directly
	fmt.Println("Passing parameters correctly:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("  Goroutine %d is running\n", id)
			time.Sleep(50 * time.Millisecond)
		}(i) // Pass 'i' as parameter
	}

	// Demonstrate closure variable capture issue
	fmt.Println("Demonstrating closure variable capture issue:")
	for i := 1; i <= 3; i++ {
		go func() {
			// This will capture the loop variable 'i' by reference
			// All goroutines will see the final value of 'i' (4)
			fmt.Printf("  Problematic goroutine sees i = %d\n", i)
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Goroutines with parameters completed!")
	fmt.Println()
}

// 3. Multiple Goroutines
// Demonstrates multiple concurrent goroutines with WaitGroup coordination
func multipleGoroutines() {
	fmt.Println("=== 3. Multiple Goroutines ===")

	var wg sync.WaitGroup

	// Create multiple goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("  Worker %d starting...\n", id)
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			fmt.Printf("  Worker %d completed\n", id)
		}(i)
	}

	fmt.Println("Waiting for all goroutines to complete...")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed!")
	fmt.Println()
}

// 4. Goroutine Communication
// Demonstrates basic producer-consumer pattern using channels
func goroutineCommunication() {
	fmt.Println("=== 4. Goroutine Communication ===")

	// Create a channel for communication
	messageChan := make(chan string, 3) // Buffered channel

	// Producer goroutine
	go func() {
		messages := []string{"Hello", "World", "from", "Go", "channels"}
		for _, msg := range messages {
			fmt.Printf("  Producer sending: %s\n", msg)
			messageChan <- msg
			time.Sleep(100 * time.Millisecond)
		}
		close(messageChan) // Close channel when done
	}()

	// Consumer goroutine
	go func() {
		for msg := range messageChan {
			fmt.Printf("  Consumer received: %s\n", msg)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("  Consumer finished (channel closed)")
	}()

	// Wait for communication to complete
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Goroutine communication completed!")
	fmt.Println()
}

// 5. Goroutine Error Handling
// Demonstrates error handling and panic recovery in goroutines
func goroutineErrorHandling() {
	fmt.Println("=== 5. Goroutine Error Handling ===")

	// Error channel for collecting errors
	errChan := make(chan error, 2)

	// Goroutine with panic recovery
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic recovered: %v", r)
			}
		}()

		fmt.Println("  Goroutine with potential panic starting...")
		time.Sleep(100 * time.Millisecond)

		// Simulate a panic
		panic("simulated panic in goroutine")
	}()

	// Goroutine that returns an error
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("  Goroutine with error starting...")
		errChan <- fmt.Errorf("simulated error in goroutine")
	}()

	// Collect errors
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case err := <-errChan:
				fmt.Printf("  Error received: %v\n", err)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("  Timeout waiting for error")
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)
	fmt.Println("Goroutine error handling completed!")
	fmt.Println()
}

// 6. Goroutine Best Practices
// Demonstrates proper cleanup, resource management, and avoiding leaks
func goroutineBestPractices() {
	fmt.Println("=== 6. Goroutine Best Practices ===")

	// Using context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Always call cancel to free resources

	// Goroutine with proper cleanup
	go func() {
		defer fmt.Println("  Goroutine cleanup completed")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("  Goroutine cancelled via context")
				return
			default:
				fmt.Println("  Goroutine working...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	// Goroutine with resource management
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("  Resource cleanup completed")

		// Simulate resource allocation
		fmt.Println("  Allocating resources...")
		time.Sleep(100 * time.Millisecond)

		// Simulate work
		fmt.Println("  Working with resources...")
		time.Sleep(100 * time.Millisecond)

		// Resources are automatically cleaned up via defer
	}()

	// Wait for goroutine to complete
	wg.Wait()

	// Wait for context cancellation
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Goroutine best practices completed!")
	fmt.Println()
}
