package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main end.")
	c := make(chan int, 3)
	quit := make(chan bool)

	go func() {
		defer fmt.Println("sub 1 goroutine end.")
		fmt.Println("start sub 1.")
		i := 0
		for {
			c <- i
			fmt.Println("sub 1 send data: ", i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		defer fmt.Println("sub 2 goroutine end.")
		fmt.Println("start sub 2.")
		time.Sleep(2 * time.Second)
		quit <- true
		fmt.Println("sub 2 send quit message.")
		close(quit)
	}()

	var n int
	n = 0
	for {
		select {
		case <-c:
			fmt.Println("[", n, "] read channel c. data: ", <-c)
			n++
		case <-quit:
			fmt.Println("recv quit")
			close(c)
			return
			//default:
			//fmt.Println("not read data.")
		}
	}
}
