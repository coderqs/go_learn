package main

import (
	"fmt"
	"time"
)

func NewTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine: i=", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go NewTask()

	i := 0
	for {
		i++
		fmt.Println("main Goroutine: i=", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("main Goroutine exit.")
}
