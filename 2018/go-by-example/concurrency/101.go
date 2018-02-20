package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	go f("goroutine2")
	time.Sleep(time.Second * 1)
}

// END OMIT
