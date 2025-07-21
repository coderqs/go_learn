package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is call...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is: ", inputType)

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is: ", inputValue)
	fmt.Println("=========================================")
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)
	}
	fmt.Println("=========================================")
	fmt.Println("inputType.NumMethod(): ", inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func main() {
	var num float64 = 1.235
	reflectNum(num)

	user := User{1, "zhang3", 23}

	DoFiledAndMethod(user)
}
