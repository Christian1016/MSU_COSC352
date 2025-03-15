package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if we have the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run hello.go <name> <number>")
		os.Exit(1)
	}

	// Retrieve the name and number from the command-line arguments
	name := os.Args[1]

	// Convert the second argument to an integer
	number, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: The second argument must be an integer.")
		os.Exit(1)
	}

	// Print the greeting the specified number of times
	for i := 0; i < number; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}
