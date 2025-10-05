package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Channel Fundamentals ===")
	fmt.Println("Demonstrating channel operations and patterns in Go")
	fmt.Println()

	// 1. Basic Channel Operations
	basicChannelOperations()

	// 2. Buffered vs Unbuffered Channels
	bufferedVsUnbufferedChannels()

	// 3. Channel Closing
	channelClosing()

	// 4. Select Statement
	selectStatement()

	// 5. Channel Patterns
	channelPatterns()

	// 6. Channel Best Practices
	channelBestPractices()
}

// 1. Basic Channel Operations
// Demonstrates channel creation, send/receive, blocking behavior, and direction
func basicChannelOperations() {
	fmt.Println("=== 1. Basic Channel Operations ===")

	// Unbuffered channel
	fmt.Println("Unbuffered channel:")
	unbuffered := make(chan string)
	go func() {
		fmt.Println("  Sending to unbuffered channel...")
		unbuffered <- "Hello from unbuffered!"
	}()
	msg := <-unbuffered
	fmt.Printf("  Received: %s\n", msg)

	// Buffered channel
	fmt.Println("Buffered channel:")
	buffered := make(chan int, 3)
	fmt.Println("  Sending to buffered channel...")
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("  All values sent (no blocking)")

	fmt.Printf("  Received: %d\n", <-buffered)
	fmt.Printf("  Received: %d\n", <-buffered)
	fmt.Printf("  Received: %d\n", <-buffered)

	// Channel direction
	fmt.Println("Channel direction:")
	sendOnly := make(chan<- string)
	receiveOnly := make(<-chan string)
	fmt.Printf("  Send-only channel: %T\n", sendOnly)
	fmt.Printf("  Receive-only channel: %T\n", receiveOnly)

	// Practical example of channel direction
	fmt.Println("Practical channel direction example:")
	ch := make(chan string)

	// Producer function - can only send
	producer := func(out chan<- string) {
		out <- "Hello"
		out <- "World"
		close(out)
		fmt.Println("  Producer finished")
	}

	// Consumer function - can only receive
	consumer := func(in <-chan string) {
		for msg := range in {
			fmt.Printf("  Consumer received: %s\n", msg)
		}
		fmt.Println("  Consumer finished")
	}

	// Start producer and consumer
	go producer(ch)
	consumer(ch)

	fmt.Println("Basic channel operations completed!")
	fmt.Println()
}

// 2. Buffered vs Unbuffered Channels
// Demonstrates blocking behavior differences and producer-consumer patterns
func bufferedVsUnbufferedChannels() {
	fmt.Println("=== 2. Buffered vs Unbuffered Channels ===")

	// Unbuffered channel blocking behavior
	fmt.Println("Unbuffered channel blocking:")
	unbuffered := make(chan int)

	// Producer goroutine
	go func() {
		fmt.Println("  Producer: sending 1...")
		unbuffered <- 1
		fmt.Println("  Producer: sent 1")
		fmt.Println("  Producer: sending 2...")
		unbuffered <- 2
		fmt.Println("  Producer: sent 2")
	}()

	// Consumer blocks until producer sends (natural synchronization)
	fmt.Printf("  Consumer: received %d\n", <-unbuffered)
	fmt.Printf("  Consumer: received %d\n", <-unbuffered)

	// Buffered channel non-blocking behavior
	fmt.Println("Buffered channel non-blocking:")
	buffered := make(chan int, 2)

	fmt.Println("  Producer: sending 1...")
	buffered <- 1
	fmt.Println("  Producer: sent 1 (no blocking)")
	fmt.Println("  Producer: sending 2...")
	buffered <- 2
	fmt.Println("  Producer: sent 2 (no blocking)")

	fmt.Printf("  Consumer: received %d\n", <-buffered)
	fmt.Printf("  Consumer: received %d\n", <-buffered)

	fmt.Println("Buffered vs unbuffered channels completed!")
	fmt.Println()
}

// 3. Channel Closing
// Demonstrates proper channel closing, detection, and graceful handling
func channelClosing() {
	fmt.Println("=== 3. Channel Closing ===")

	// Channel closing and detection
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("  Sent: %d\n", i)
		}
		close(ch)
		fmt.Println("  Channel closed")
	}()

	// Method 1: Using range
	fmt.Println("Receiving with range:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}

	// Method 2: Manual detection
	fmt.Println("Manual closed channel detection:")
	ch2 := make(chan string)
	go func() {
		ch2 <- "Hello"
		close(ch2)
	}()

	value, ok := <-ch2
	if ok {
		fmt.Printf("  Received: %s (channel open)\n", value)
	}

	_, ok = <-ch2
	if !ok {
		fmt.Println("  Channel is closed")
	}

	fmt.Println("Channel closing completed!")
	fmt.Println()
}

// 4. Select Statement
// Demonstrates select for multiple channel operations, timeouts, and default cases
func selectStatement() {
	fmt.Println("=== 4. Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Multiple channel operations
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	fmt.Println("Select with multiple channels:")
	select {
	case msg := <-ch1:
		fmt.Printf("  Received from ch1: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("  Received from ch2: %s\n", msg)
	}

	// Timeout handling
	fmt.Println("Select with timeout:")
	timeoutCh := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutCh <- "too late"
	}()

	select {
	case msg := <-timeoutCh:
		fmt.Printf("  Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  Timeout occurred")
	}

	// Default case (non-blocking)
	fmt.Println("Select with default (non-blocking):")
	nonBlockingCh := make(chan string)

	select {
	case msg := <-nonBlockingCh:
		fmt.Printf("  Received: %s\n", msg)
	default:
		fmt.Println("  No message available (non-blocking)")
	}

	fmt.Println("Select statement completed!")
	fmt.Println()
}

// 5. Channel Patterns
// Demonstrates fan-in, fan-out, pipelines, and generator patterns
func channelPatterns() {
	fmt.Println("=== 5. Channel Patterns ===")

	// Fan-out pattern
	fmt.Println("Fan-out pattern:")
	input := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	// Create multiple workers
	worker1 := fanOut(input)
	worker2 := fanOut(input)

	// Fan-in pattern
	fmt.Println("Fan-in pattern:")
	output := fanIn(worker1, worker2)

	for value := range output {
		fmt.Printf("  Received: %d\n", value)
	}

	// Generator pattern
	fmt.Println("Generator pattern:")
	gen := generator(1, 2, 3, 4, 5)
	for i := 0; i < 5; i++ {
		fmt.Printf("  Generated: %d\n", <-gen)
	}

	// Pipeline pattern
	fmt.Println("Pipeline pattern:")
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()

	for square := range squares {
		fmt.Printf("  Square: %d\n", square)
	}

	fmt.Println("Channel patterns completed!")
	fmt.Println()
}

// Fan-out worker
func fanOut(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			output <- n * 2 // Double the value
		}
	}()
	return output
}

// Fan-in multiplexer
func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				output <- n
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Generator function
func generator(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}

// 6. Channel Best Practices
// Demonstrates avoiding leaks, proper error handling, and channel ownership
func channelBestPractices() {
	fmt.Println("=== 6. Channel Best Practices ===")

	// Proper channel ownership
	fmt.Println("Channel ownership:")
	producer := func() <-chan string {
		ch := make(chan string)
		go func() {
			defer close(ch)
			for i := 1; i <= 3; i++ {
				ch <- fmt.Sprintf("message-%d", i)
			}
		}()
		return ch
	}

	consumer := func(input <-chan string) {
		for msg := range input {
			fmt.Printf("  Consumed: %s\n", msg)
		}
	}

	ch := producer()
	consumer(ch)

	// Error handling with channels
	fmt.Println("Error handling:")
	errCh := make(chan error, 1)
	resultCh := make(chan string, 1)

	go func() {
		defer close(errCh)
		defer close(resultCh)

		// Simulate work
		time.Sleep(50 * time.Millisecond)

		// Simulate error
		errCh <- fmt.Errorf("simulated error")
	}()

	select {
	case err := <-errCh:
		fmt.Printf("  Error: %v\n", err)
	case result := <-resultCh:
		fmt.Printf("  Result: %s\n", result)
	}

	// Avoiding channel leaks
	fmt.Println("Avoiding channel leaks:")
	leakyCh := make(chan int)
	go func() {
		// This goroutine will leak if no one reads from the channel
		leakyCh <- 42
		fmt.Println("  This won't print (goroutine blocked)")
	}()

	// Proper cleanup
	go func() {
		time.Sleep(100 * time.Millisecond)
		<-leakyCh // Clean up the leaked goroutine
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Channel best practices completed!")
	fmt.Println()
}
