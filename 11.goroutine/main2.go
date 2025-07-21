package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("defer main A")

	go func(a int, b int) {
		defer fmt.Println("defer goroutine B")
		go func() {
			defer fmt.Println("defer goroutine C")
			runtime.Goexit()
			fmt.Println("hello goroutine 3")

		}()
		//i := 0
		//for {
		//	i++
		//	fmt.Println("goroutine i = ", i)
		//	time.Sleep(1 * time.Second)
		//}
	}(1, 2)
	time.Sleep(1 * time.Second)
	//i := 0
	//for {
	//	i++
	//	fmt.Println("main goroutine i = ", i)
	//	time.Sleep(1 * time.Second)
	//}
}
