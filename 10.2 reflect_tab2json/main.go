package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name   string   `json:"name"`
	year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2008, 30, []string{"zhouxingxing", "zhangbaizhi"}}
	// 结构体 --> json
	jsonStr, ec := json.Marshal(movie)
	if ec != nil {
		fmt.Println("json marshal error.")
		return
	}
	fmt.Printf("jsonStr: %s\n", jsonStr)

	// json --> 结构体
	new_movie := Movie{}
	err := json.Unmarshal(jsonStr, &new_movie)
	if err != nil {
		fmt.Println("json unmarshal error.")
		return
	}
	fmt.Printf("%v\n", new_movie)
}
