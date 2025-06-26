package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk()...")
}

type SuperMan struct {
	Human // 集成 Human 的方法
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()...")
}

func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	s := SuperMan{Human{"li4", "female"}, 100}
	s.Walk()
	s.Eat()
	s.Fly()

	var s1 SuperMan
	s1.name = "wang5"
	s1.sex = "female"
	s1.level = 88

}
