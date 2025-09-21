package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// === Context Package ===
	// This module demonstrates context usage for cancellation and timeouts in Go.
	// Complete the following exercises:

	// TODO: Implement the following context patterns:

	// 1. Basic Context Operations
	//    - Create context.Background()
	//    - Use context.TODO() for placeholders
	//    - Understand context hierarchy
	//    - Implement context propagation

	// 2. Context with Timeout
	//    - Use context.WithTimeout()
	//    - Handle timeout scenarios
	//    - Implement cascading timeouts
	//    - Handle timeout errors

	// 3. Context with Cancellation
	//    - Use context.WithCancel()
	//    - Implement manual cancellation
	//    - Handle cancellation propagation
	//    - Implement cleanup on cancellation

	// 4. Context with Deadline
	//    - Use context.WithDeadline()
	//    - Handle deadline scenarios
	//    - Implement deadline-based operations
	//    - Handle deadline exceeded errors

	// 5. Context with Values
	//    - Use context.WithValue()
	//    - Implement request-scoped data
	//    - Handle context value propagation
	//    - Avoid context value abuse

	// 6. Context Best Practices
	//    - Pass context as first parameter
	//    - Don't store context in structs
	//    - Use context for cancellation only
	//    - Implement proper context cleanup

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining concepts
	// - Test with different scenarios

	fmt.Println("=== Context Package ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each context pattern!")
	fmt.Println()

	// Example implementations (to be replaced with your code):
	fmt.Println("Example implementations:")

	// Example 1: Context with timeout
	fmt.Println("\n1. Context with Timeout:")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Work completed successfully")
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
		}
	}()

	time.Sleep(3 * time.Second)

	// Example 2: Context with cancellation
	fmt.Println("\n2. Context with Cancellation:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Goroutine cancelled:", ctx2.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	cancel2()
	time.Sleep(100 * time.Millisecond)

	// Example 3: Context with deadline
	fmt.Println("\n3. Context with Deadline:")
	deadline := time.Now().Add(1 * time.Second)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()

	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("This should not print")
		case <-ctx3.Done():
			fmt.Println("Deadline reached:", ctx3.Err())
		}
	}()

	time.Sleep(2 * time.Second)
}
