package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"my name"`
	Sex  string `info:"sex" doc:"my sex"` // : 后面不能有空格，有空格会将value识别为空
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo)
		fmt.Println("doc: ", tagdoc)
	}
}

func main() {
	u := resume{"li4", "man"}
	findTag(&u)
}
