package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() { // main 函数只能在 main package 包里，且不能有任何参数和返回值
	x, y := "str_x", "str_y"
	x1, y1 := swap(x, y)
	fmt.Println("x=", x, ", y=", y)
	fmt.Println("x1=", x1, ", y1=", y1)
}
