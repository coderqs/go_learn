package main

import "fmt"

func main() {
	defer fmt.Println("main goroutine.")
	c := make(chan int)
	go func() {
		defer fmt.Println("sub goroutine.")
		fmt.Println("send to main goroutine 666")
		c <- 666
	}()

	num := <-c

	fmt.Println("num = ", num)
	fmt.Println("main end.")
}
