package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main goroutine end.")
	c := make(chan int, 4)

	go func() {
		defer fmt.Println("sub goroutine end.")

		for i := 0; i < 4; i++ {
			c <- i
		}

		close(c)
	}()

	time.Sleep(1 * time.Second)

	go func() {
		defer fmt.Println("sub 2 goroutine end.")
		for {
			if num, ok := <-c; ok {
				fmt.Println("sub2, num = ", num)
			} else {
				break
			}
		}
	}()

	//for {
	//	if num, ok := <-c; ok {
	//		fmt.Println("num = ", num)
	//	} else {
	//		break
	//	}
	//}
	for num := range c {
		fmt.Println("main num = ", num)
	}

	fmt.Println("main end.")
}
