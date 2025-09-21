package main

import (
	"fmt"
	"time"
)

func main() {
	// === Channel Fundamentals ===
	// This module demonstrates basic channel operations and patterns in Go.
	// Complete the following exercises:

	// TODO: Implement the following channel patterns:

	// 1. Basic Channel Operations
	//    - Create channels (buffered and unbuffered)
	//    - Send and receive data
	//    - Understand channel blocking behavior
	//    - Handle channel direction (send-only, receive-only)

	// 2. Buffered vs Unbuffered Channels
	//    - Compare buffered and unbuffered channels
	//    - Understand blocking behavior differences
	//    - Implement producer-consumer patterns
	//    - Handle channel capacity

	// 3. Channel Closing
	//    - Properly close channels
	//    - Detect closed channels
	//    - Handle closed channel gracefully
	//    - Use range with channels

	// 4. Select Statement
	//    - Use select for multiple channel operations
	//    - Implement timeout handling
	//    - Handle default case
	//    - Implement non-blocking operations

	// 5. Channel Patterns
	//    - Implement fan-in pattern
	//    - Implement fan-out pattern
	//    - Create channel pipelines
	//    - Implement generator pattern

	// 6. Channel Best Practices
	//    - Avoid channel leaks
	//    - Implement proper error handling
	//    - Use channels for communication
	//    - Handle channel ownership

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining concepts
	// - Test with different scenarios

	fmt.Println("=== Channel Fundamentals ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each channel pattern!")
	fmt.Println()

	// Example implementations (to be replaced with your code):
	fmt.Println("Example implementations:")

	// Example 1: Basic channel
	fmt.Println("\n1. Basic Channel Example:")
	channel := make(chan string)
	go func() {
		channel <- "Hello, World!"
	}()
	message := <-channel
	fmt.Println(message)

	// Example 2: Buffered channel
	fmt.Println("\n2. Buffered Channel Example:")
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println("Sent to buffered channel")
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)

	// Example 3: Select statement
	fmt.Println("\n3. Select Statement Example:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout!")
	}
}
