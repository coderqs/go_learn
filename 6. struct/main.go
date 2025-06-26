package main

import "fmt"

type myint int // 别名，类似 C++ typedef

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 值传递
	book.auth = "66"
}

func changeBook2(book *Book) {
	// 指针传递
	book.auth = "77"
}

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)

}

func (this Hero) GetName() string {
	return this.Name
}

// func (this Hero) SetName(name string) {
// 	this.Name = name
// }
func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	var a myint = 10
	fmt.Printf("a = %d, a type: %T\n", a, a)

	var book1 Book
	book1.title = "golong"
	book1.auth = "zhangsan"

	fmt.Printf("0. book1 val: %v, type: %T\n", book1, book1)

	changeBook(book1)

	fmt.Printf("1. book1 val: %v, type: %T\n", book1, book1)

	changeBook2(&book1)

	fmt.Printf("2. book1 val: %v, type: %T\n", book1, book1)

	////////////////////////////////
	var my_hero Hero
	my_hero.Name = "zhang3"
	my_hero.Ad = 24
	my_hero.Level = 10

	my_hero.Show()
	my_hero.SetName("li4")
	my_hero.Show()
}
