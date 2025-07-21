package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine end.")
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("sub goroutine end.")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("i = ", i, ", len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("recv num = ", num)
	}

	fmt.Println("main end.")
}
