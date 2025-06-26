package main

import "fmt"

var a int             // 未初始化状态下，默认值 0
var b = "string_type" // 根据赋值自动判断类型
var e, f int          // 多变量定义
var h, g = 11, "str1"

const (
	EC_FAILED    = -1
	EC_SUCCEED   = 0
	EC_NOT_FOUND = 2
)

const (
	EC_MY_FAILED = iota - 1
	EC_MY_SUCCEED
	EC_MY_NOT_FOUND
)
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := 10 // 不写 var 直接赋值，仅在函数体中可用
	fmt.Printf("c=%d", c)
	fmt.Println("e=", e, "f=", f, "h=", h, "g=", g)

	_, value := 22, 43
	fmt.Println("value=", value)

	fmt.Println("============================================")
	fmt.Println("EC_FAILED =", EC_FAILED,
		"EC_SUCCEED =", EC_SUCCEED,
		"EC_NOT_FOUND =", EC_NOT_FOUND)
	fmt.Println("EC_MY_FAILED  = ", EC_MY_FAILED,
		"EC_MY_SUCCEED = ", EC_MY_SUCCEED,
		"EC_MY_NOT_FOUND = ", EC_MY_NOT_FOUND)

	fmt.Println("Apple=", Apple, ", Banana=", Banana,
		"Cherimoya=", Cherimoya, ", Durian=", Durian,
		"Elderberry=", Elderberry, ", Fig=", Fig)
}
