package main

import (
	"fmt"
	"time"
)

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs, results := make(chan int), make(chan int)
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}

// END OMIT
