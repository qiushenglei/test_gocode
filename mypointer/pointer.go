package mypointer

import (
	"fmt"
	"reflect"
)

type Astruct struct {
	name string
}

func AssignInterface() {
	var a interface{}

	b := Astruct{
		name: "hi",
	}
	var c Astruct

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
	fmt.Printf("值a=%s 值b=%s 值c=%s\n", a, b, c)
	fmt.Printf("地址a=%p 地址b=%p 地址c=%p\n", &a, &b, &c)
	fmt.Println("a:", &a, "b:", &b, "c:", &c)
	a = &b
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
	fmt.Printf("值a=%s 值b=%s 值c=%s\n", a, b, c)
	fmt.Printf("地址a=%p 地址b=%p 地址c=%p\n", a, &b, &c)
	fmt.Println("a:", &a, "b:", &b, "c:", &c)

	var e *int
	fmt.Println(reflect.TypeOf(e))
	fmt.Printf("地址%p\n", e)
	e = new(int)
	fmt.Printf("地址%p\n", e)

	f := 3
	fmt.Printf("地址%p\n", f)
}

func TestPo() {
	num := 5

	fmt.Println(num)
	fmt.Println(&num) //打印num的地址

	fmt.Println("----------from")
	from := &num
	fmt.Println(reflect.TypeOf(from))
	fmt.Println(from)               //打印from保存的地址(就是num变量的地址)
	fmt.Println(*from)              //打印from保存的地址的值（就是num的值）
	fmt.Printf("from地址%p\n", from)  //打印from值保存的地址
	fmt.Printf("from地址%p\n", &from) //打印from变量的地址

	fmt.Println("----------from1")
	from1 := new(int) //创建*int类型的from1变量，当前from内存里保存的值是0x0
	*from1 = num      //from1的值，指向5，并不是指向num。
	fmt.Println(reflect.TypeOf(from1))
	fmt.Println(from1)           // 打印from1保存的地址，这样可以发现，打印的并不是num字段你的地址
	fmt.Println(*from1)          // 打印from1字段地址
	fmt.Printf("地址%p\n", from1)  // 打印from1值保存的的地址
	fmt.Printf("地址%p\n", &from1) // 打印from1变量的的地址
}

type A struct {
	A int
}

func TestVarmake() {
	var a int
	var b []int
	e := make([]int, 1)
	var c A
	var d *A

	f := new(int)
	g := new(A)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", b)

	fmt.Printf("e:%p\n", e)
	fmt.Printf("c:%p\n", &c)
	fmt.Printf("%p\n", d)
	fmt.Println(*f)
	fmt.Println(*g)

}
