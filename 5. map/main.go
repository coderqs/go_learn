package main

import "fmt"

func main() {
	var my_map map[string]string
	if my_map == nil {
		fmt.Println("my_map is empty.")
	}

	// 使用map 前需要先用 make 给 map 分配空间
	my_map = make(map[string]string, 10)
	my_map["one"] = "golong"
	my_map["two"] = "C/C++"
	my_map["three"] = "python"
	fmt.Println(my_map)

	// 第二种声明方式
	my_map2 := make(map[string]string)
	my_map2["one"] = "golong+s"
	my_map2["two"] = "C/C++ 20"
	my_map2["three"] = "python3"
	fmt.Println(my_map2)

	// 第三种方式
	my_map3 := map[string]string{
		"one":   "php",
		"two":   "rules",
		"three": "java",
	}
	fmt.Println(my_map3)

	////////////////////////////////////////
	city_map := map[string]string{
		"china": "beijing",
		"japan": "tokyo",
		"usa":   "newyork",
	}
	// 遍历
	for key, val := range city_map {
		fmt.Println("key: ", key)
		fmt.Println("val: ", val)
	}
	// 删除
	delete(city_map, "japan")
	fmt.Println("1. ", city_map)
	delete(city_map, "newyork") //使用值不能删除，只能使用key删除
	fmt.Println("2. ", city_map)
	// 修改
	city_map["china"] = "BEIJING"
	fmt.Println("3. ", city_map)

}
