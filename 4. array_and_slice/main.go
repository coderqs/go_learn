package main

import "fmt"

// 在函数做形参的时候，数组是值传递，动态数组是引用传递，并且数组只能接收和自己大小一致的数组。

func printArray(arr []int) {
	for _, val := range arr {
		fmt.Println("val = ", val)
	}

	arr[0] = 100
}

func main() {
	var array_1 [10]int

	for i := 0; i < len(array_1); i++ {
		fmt.Println(array_1[i])
	}

	array_2 := [10]int{1, 2, 3, 4}
	for index, val := range array_2 {
		fmt.Println("[", index, "] ", val)
	}

	array_3 := []int{1, 2, 3, 4} // 动态数组 切片 slice
	fmt.Printf("array_3 type is %T\n", array_3)
	printArray(array_3)
	for _, val := range array_3 {
		fmt.Println("array_3 val = ", val)
	}
	//--------------------------------------
	fmt.Println("===============================================")
	// 声明 slice1 是一个切片，并且初始化，默认值是1
	slice1 := []int{1, 2, 3}
	fmt.Printf("len=%d, slice1=%v\n", len(slice1), slice1)

	// 声明 slice2 是一个切片，但是没有给 slice 分配空间
	var slice2 []int
	slice2 = make([]int, 3) // 开辟3个空间，默认值为0
	fmt.Printf("len=%d, slice2=%v\n", len(slice2), slice2)

	// 声明 slice3 是一个切片，同时给slice分配空间，开辟3个空间，默认值为0
	var slice3 []int = make([]int, 3)
	fmt.Printf("len=%d, slice3=%v\n", len(slice3), slice3)

	// 声明 slice4 是一个切片，同时给slice分配空间，开辟3个空间，默认值为0 通过 := 推导出 slice4 是一个切片
	slice4 := make([]int, 3)
	fmt.Printf("len=%d, slice4=%v\n", len(slice4), slice4)

	// 判断 silce 是否为空
	if slice1 == nil {
		fmt.Println("slice1 is empty.")
	} else {
		fmt.Println("slice1 not empty.")
	}

	//--------------------------------------
	fmt.Println("===============================================")
	var numbers []int = make([]int, 3, 5)
	fmt.Printf("len: %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers追加一个元素1，numbers len = 4 [0,0,0,0]， cap=5
	numbers = append(numbers, 1)
	fmt.Printf("len: %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 当 append 追加元素超过容量的大小，则会自动扩大原容量的一倍
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Printf("len: %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//--------------------------------------
	fmt.Println("===============================================")
	// 切片的截取，直接赋值的话是浅拷贝
	s := []int{1, 2, 3}
	s1 := s[0:2] // [1,2]
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println("s = ", s, "s1 = ", s1)

	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println("s2 = ", s2)

	s1[1] = 200
	fmt.Println("s = ", s, "s1 = ", s1, "s2 = ", s2)

}
