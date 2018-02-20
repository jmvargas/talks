package main

import (
	"fmt"
)

func main() {
	go fmt.Println("World")
	fmt.Println("Hello")
}
