package main

import "fmt"

// interface{} 是万能数据类型，可以承接任何类型，有点类似 C++ 中的 void
func PrintData(arg interface{}) {
	fmt.Println("PrintData param: ", arg)

	// interface{} 如何区分，当前引用的底层数据类型是什么
	// interface{} 提供了“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type")
	}
	fmt.Printf("arg type is: %T,\n", value)
	fmt.Println("arg value: ", value)
}

type Book struct {
	name string
	auth string
}

func main() {
	book := Book{"Golang learn", "qs"}

	PrintData(book)
	PrintData(100)
	PrintData("hello")
	PrintData(3.14)
}
