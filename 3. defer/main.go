package main

import "fmt"

// defer 类似于C++中的析构函数
// defer 是压栈执行的，所以写前面的后执行
// defer 在 return 之后调用，因为代码是在 } 之后才开始出栈

func func1() {
	fmt.Println("my func1")
}

func func2() {
	fmt.Println("my func2")
}

func func3() {
	fmt.Println("my func3")
}

func funcDefer() int {
	fmt.Println("my defer")
	return 0
}

func funcReturn() int {
	fmt.Println("my return")
	return 0
}

func funcDeferAndReturn() int {
	defer funcDefer()

	return funcReturn()
}

func main() {
	//defer fmt.Println("main end")
	//defer fmt.Println("main defer")
	//defer func1()
	//defer func2()
	//defer func3()

	//fmt.Println("main: print 1")
	//fmt.Println("main: print 2")

	funcDeferAndReturn()
}
