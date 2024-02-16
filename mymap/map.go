package mymap

import (
	"fmt"
	"time"
)

func updateTask1(m map[string]int) {
	m["key"] = 100
}

func updateTask2(m map[string]int) {
	m["key"] = 200
}

func updateTask3(m map[string]int) {
	m["key"] = 300
}

func UpdateMap() {
	myMap := make(map[string]int)

	myMap["key"] = 10

	fmt.Println("Before:", myMap)
	fmt.Printf("Before memery: %p\n", myMap)

	updateTask1(myMap)

	fmt.Println("After:", myMap)

	go updateTask2(myMap)

	go updateTask3(myMap)

	time.Sleep(time.Second * 2)
	fmt.Println("After:", &myMap)
	fmt.Printf("After memery: %p\n", myMap)
}

func Cap() {
	a := make(map[string]interface{}, 1)

	a["as"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))
	a["as1"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))

	a["as2"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))

	a["as3"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))
	a["as4"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))
	a["as5"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))
	a["as6"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))
	a["as7"] = 1233
	fmt.Printf("%p %d\n ", a, len(a))

	fmt.Println(len(a))
}

func Copy() {
	// map不能copy
	//a := map[string]interface{}{
	//	"a": 1,
	//	"b": "asd",
	//}
	//b := make(map[string]interface{})
	//copy(b, a)

	// slice copy
	a := make([]int, 0)

	a = append(a, 1123)
	b := make([]int, len(a))
	//var b []int
	copy(a, b)
	fmt.Printf("%p %v\n ", a, len(a))
	fmt.Printf("%p %v\n ", b, len(b))
}

func Cap1() {
	a := make([]int, 0)

	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))
	a = append(a, 1123)
	fmt.Printf("%p %v\n ", a, cap(a))

	fmt.Println(len(a))
}

func Delete() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	delete(m, "banana")
	fmt.Println(m)
}

func ThreadSafe() {
	//sync.Map{}
}

type AA struct {
	data []int
}

func TestNil() (res *AA) {
	var a map[int]string
	fmt.Println(len(a))
	return
}

func TestNil1() (res map[int]int) {

	return
}

func TestSNil() {
	res := TestNil1()
	fmt.Println()
}
