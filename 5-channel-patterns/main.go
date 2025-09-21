package main

import (
	"fmt"
)

func main() {
	// === Channel Patterns ===
	// This module demonstrates advanced channel patterns in Go.
	// Complete the following exercises:

	// TODO: Implement the following patterns:

	// 1. Select Statement Pattern
	//    - Create multiple channels
	//    - Use select to handle multiple channel operations
	//    - Implement timeout handling
	//    - Handle default case

	// 2. Channel Closing Pattern
	//    - Demonstrate proper channel closing
	//    - Detect closed channels
	//    - Handle closed channel gracefully

	// 3. Pipeline Pattern
	//    - Create a 3-stage pipeline
	//    - Stage 1: Generate numbers
	//    - Stage 2: Square the numbers
	//    - Stage 3: Print the results

	// 4. Fan-in Pattern
	//    - Create multiple producer goroutines
	//    - Merge all outputs into a single channel
	//    - Use select to multiplex channels

	// 5. Fan-out Pattern
	//    - Create multiple consumer goroutines
	//    - Distribute work from a single channel
	//    - Implement load balancing

	// 6. Generator Pattern
	//    - Create a function that returns a channel
	//    - Generate values on demand
	//    - Implement cleanup when done

	// Instructions:
	// - Implement each pattern in separate functions
	// - Add proper error handling
	// - Use meaningful variable names
	// - Add comments explaining the pattern
	// - Test each pattern with different scenarios

	fmt.Println("=== Channel Patterns ===")
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each pattern one by one!")
}
