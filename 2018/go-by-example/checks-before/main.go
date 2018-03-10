package main

import (
	"fmt"
)

func main() {
	// Compilation fails if we have unused variables
	a := "Not used"
	fmt.Printf("Hello!")
}
