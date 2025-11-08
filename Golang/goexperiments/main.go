package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("entered in worker func with id", id)
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	fmt.Println("Closing jobs channel")
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		tempres := <-results
		fmt.Println("result", tempres)
	}
	fmt.Println("Completed function")
}
