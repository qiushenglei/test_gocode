package myMemory

import "fmt"

func Memory() {
	var a *int
	b := 1
	a = &b
	fmt.Println(a)
	fmt.Println(b)

	c := 1
	d := &c
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(*d)
}

func Memory1() {
	b := 1
	a := new(int)
	*a = b
	fmt.Println(a)
}

type Person struct {
	Name string
	Age  int
}

func NewTest() {
	// 分配一个 Person 结构体在堆上，并返回指向该结构体的指针
	p := new(Person)

	//a := Person{}

	fmt.Println(p.Name)
}
