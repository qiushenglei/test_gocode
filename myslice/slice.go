package myslice

import "fmt"

func testSlice1() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sli := arr[3:]
	testSlice2(sli)
	fmt.Println("原数组", arr, "原数组len", len(arr), "原数组cap", cap(arr))
	fmt.Println("原切片", sli, "原切片len", len(sli), "原切片cap", cap(sli))
}

func testSlice2(sli []int) {
	sli1 := append(sli, 10)
	sli1[0] = 999
	fmt.Println(sli1, sli)
}

func testSlice() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli := arr[3:8]
	fmt.Println("原数组", arr, "原数组len", len(arr), "原数组cap", cap(arr))
	fmt.Println("原切片", sli, "原切片len", len(sli), "原切片cap", cap(sli))
	sli1 := append(sli, 10)
	fmt.Println("原切片append10后")
	fmt.Println("原数组", arr, "原数组len", len(arr), "原数组cap", cap(arr))
	fmt.Println("原切片", sli, "原切片len", len(sli), "原切片cap", cap(sli))
	fmt.Println("sli1切片", sli1, "sli1切片len", len(sli1), "sli1切片cap", cap(sli1))
	sli2 := arr[3:10]
	fmt.Println("重新在原数组重切片")
	fmt.Println("sli2切片", sli2, "sli2切片len", len(sli2), "sli2切片cap", cap(sli2))
	sli3 := append(sli2, 11)
	fmt.Println("原切片append11后")
	fmt.Println("原数组", arr, "原数组len", len(arr), "原数组cap", cap(arr))
	fmt.Println("sli3切片", sli3, "sli3切片len", len(sli3), "sli3切片cap", cap(sli3))
}

func SliceCopy() {

	sli := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli1 := sli
	var sli2 = make([]int, 10)
	copy(sli2, sli)
	sli[1] = 99
	fmt.Println(sli, sli1, sli2)
}

func SliceMakeSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len=6 cap=6 [2 3 5 7 11 13]

	// 截取切片使其长度为 0
	// pointer指向了数组的头部2
	// len是可读写的是0个元素
	// cap是pointer到数组尾部的长度，所以是6
	s1 := s[:0]
	//fmt.Println(s1[5])
	printSlice(s1) //len=0 cap=6 []

	s2 := s1[1:3]
	printSlice(s2) //len=0 cap=6 []
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func InitSlice() {
	fmt.Println("\n\n")
	var a []int
	c := []int{1, 2, 3, 4}
	b := make([]int, 0, 10)
	fmt.Printf("地址%p\n", a)
	fmt.Printf("地址%p\n", b)
	fmt.Printf("地址%p\n", &b)
	b = c[1:2]
	fmt.Printf("地址%p\n", b)
	fmt.Printf("地址%p\n", &b)
	fmt.Println(len(b), cap(b))
}

func InitSlice1() {
	var a [3]int
	var b []int
	a[0] = 3
	fmt.Println(a)
	b = append(b, 4)
	fmt.Println(b)
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

func IsParam() {
	// slice copy
	a := make([]int, 0, 2)

	a = append(a, 1123)
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p\n", a)
	fmt.Println(a)
	Param(a)
	fmt.Println(len(a), cap(a))
	fmt.Printf("%p\n", a)
	fmt.Println(a)
	fmt.Println(&a[1])

}

func Param(a []int) {
	fmt.Printf(" before %p\n", a)
	fmt.Println(&a[0])
	a = append(a, 3)
	fmt.Println(&a[1])

	//a[0] = 3
	fmt.Println(len(a), cap(a))
	fmt.Println(&a[0])
	fmt.Printf(" after %p\n", a)
	fmt.Println(a)
}

func ArrToSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]                  //2，3，4  cap = 2~9 = 8
	fmt.Println(s1, len(s1), cap(s1)) // 3， 8

	s2 := s1[2:6:7]                   //4,5,6,7 cap = 5 为什么等于5，s1的容量是2~9，那s2的容量是s1的开头到s1的第7个所以底层数组是[2,3,4,5,6,7,8]
	fmt.Println(s2, len(s2), cap(s2)) // 4， 5

	s2 = append(s2, 100) // 4,5,6,7,100
	s2 = append(s2, 200) // 4,5,6,7,100,200 超出cap，创建一个底层数组

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	fmt.Printf("%p\n", s)
	s = append(s, 100)
	fmt.Printf("%p\n", s)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func IsParam1() {
	s := []int{1, 1, 1}
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%p\n", s)
	newS := myAppend(s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
	fmt.Println(newS)
}

func Delete() {
	s := make([]int, 0, 10)
	fmt.Printf("%p\n", s)
	//s = []int{0, 1, 2, 33, 4, 5, 6, 7, 8, 9}
	//fmt.Println(&s[0])
	s = append(s, 10)
	fmt.Println(&s[0])
}

type SlicetStruct []int

func SlicetStructFunc() {
	a := SlicetStruct{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make(SlicetStruct, 3)
	fmt.Println(a, b)
}
