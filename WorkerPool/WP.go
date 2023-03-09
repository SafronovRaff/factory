package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	const jobsCount, workerCount = 15, 15
	jobs := make(chan int, jobsCount)
	result := make(chan int, jobsCount)

	for i := 0; i < workerCount; i++ {

		go worker(1, jobs, result)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("resuld %d: value = %d \n", i+1, <-result)
	}
	fmt.Println("time elapsed:", time.Since(t).String())
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)

		fmt.Printf("worker #%d finished\n", id)
		result <- j * j
	}

}
