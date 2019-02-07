package main

import (
	"fmt"
)


func main()  {
	fmt.Println("Hello word!")

	jobs := make(chan int, 25)
	results := make(chan int, 25)

	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 24; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 24; j++ {
		fmt.Println(<-results)
	}

}


func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}