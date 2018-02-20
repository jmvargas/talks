package main

import (
	"fmt"
)

func main() {
	// Compilation fails if we have unused variables
	a := 5
	fmt.Printf("Hello!")
}
