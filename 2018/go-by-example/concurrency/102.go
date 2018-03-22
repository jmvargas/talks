package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	go fmt.Println("World")
}
