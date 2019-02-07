package main

import (
	"fmt"
	//"time"
)

func say(a []int, ch chan int) {
		var sum int = 0

		for _, value := range a {
			sum += value
		}

		ch <- sum
}

func count(s int, ch chan int) {
	for i := 0; i <= s; i++ {
		ch <- i
	}
	close(ch)
}

func main()  {

	a := []int{1, 2, 3, 4, 5, 6}
	first_half := a[len(a)/2:]
	second_half := a[:len(a)/2]

	ch := make(chan int)


	go say(first_half, ch)
	go say(second_half, ch)
	
	first_half_sum, second_half_sum := <- ch, <- ch
	fmt.Println(first_half_sum)
	fmt.Println(second_half_sum)

	c := make(chan int)

	go count(5, c)

	for msg := range c {
		print(msg)
	}

	/*
	ch1 := make(chan string)
	ch2 := make(chan string)


	go func() {
		for {
			ch1 <- "Every 500 Ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "Every 2 S"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case <-ch1:
			fmt.Println(<-ch1)
		case <-ch2:
			fmt.Println(<-ch2)
		}
	}
	*/
	
}