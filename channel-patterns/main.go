package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Channel Patterns Examples ===")
	fmt.Println("This module demonstrates advanced channel patterns in Go.")
	fmt.Println("Complete the following exercises:")
	fmt.Println()

	// TODO: Implement the following patterns:

	// 1. Select Statement Pattern
	fmt.Println("1. SELECT STATEMENT PATTERN")
	fmt.Println("   - Create multiple channels")
	fmt.Println("   - Use select to handle multiple channel operations")
	fmt.Println("   - Implement timeout handling")
	fmt.Println("   - Handle default case")
	fmt.Println()

	// 2. Channel Closing Pattern
	fmt.Println("2. CHANNEL CLOSING PATTERN")
	fmt.Println("   - Demonstrate proper channel closing")
	fmt.Println("   - Detect closed channels")
	fmt.Println("   - Handle closed channel gracefully")
	fmt.Println()

	// 3. Pipeline Pattern
	fmt.Println("3. PIPELINE PATTERN")
	fmt.Println("   - Create a 3-stage pipeline")
	fmt.Println("   - Stage 1: Generate numbers")
	fmt.Println("   - Stage 2: Square the numbers")
	fmt.Println("   - Stage 3: Print the results")
	fmt.Println()

	// 4. Fan-in Pattern
	fmt.Println("4. FAN-IN PATTERN")
	fmt.Println("   - Create multiple producer goroutines")
	fmt.Println("   - Merge all outputs into a single channel")
	fmt.Println("   - Use select to multiplex channels")
	fmt.Println()

	// 5. Fan-out Pattern
	fmt.Println("5. FAN-OUT PATTERN")
	fmt.Println("   - Create multiple consumer goroutines")
	fmt.Println("   - Distribute work from a single channel")
	fmt.Println("   - Implement load balancing")
	fmt.Println()

	// 6. Generator Pattern
	fmt.Println("6. GENERATOR PATTERN")
	fmt.Println("   - Create a function that returns a channel")
	fmt.Println("   - Generate values on demand")
	fmt.Println("   - Implement cleanup when done")
	fmt.Println()

	fmt.Println("Instructions:")
	fmt.Println("- Implement each pattern in separate functions")
	fmt.Println("- Add proper error handling")
	fmt.Println("- Use meaningful variable names")
	fmt.Println("- Add comments explaining the pattern")
	fmt.Println("- Test each pattern with different scenarios")
	fmt.Println()
	fmt.Println("Run: go run main.go")
	fmt.Println("Then implement each pattern one by one!")
}
