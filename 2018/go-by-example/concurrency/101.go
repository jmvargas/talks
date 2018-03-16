package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	f("direct")
	go f("goroutine1")
	go f("goroutine2")
	time.Sleep(time.Second * 1)
}
