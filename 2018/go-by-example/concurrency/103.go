package main

import (
	"fmt"
	"time"
)

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	jobs, results := make(chan int), make(chan int)
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}

// END OMIT
