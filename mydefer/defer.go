package mydefer

//	func A(b int) {
//		fmt.Println("A func print: ", b, &b)
//	}
//
//	func JubuChuanCan() {
//		a, b := 1, 2
//
//		defer func(b int) {
//			fmt.Println("Lie func print: ", b, &b, a, &a)
//		}(b)
//
//		defer A(b)
//		fmt.Println("JubuChuanCan func print: ", b, &b, a, &a)
//		a++
//		b++
//	}
//
//	func Example() int {
//		defer fmt.Println("Deferred statement1")
//
//		num := 10
//		defer fmt.Println("Deferred statement2:", num) // 打印10，可以理解为值拷贝
//
//		defer func() {
//			fmt.Println("Deferred func statement3:", num) // 打印20,按num逃逸来理解是对的，所以是外面num的指针
//		}()
//
//		num = 20
//		return num
//	}
var num = 3

//func Example1() int {
//	defer fmt.Println("Deferred statement")
//
//	num = 10
//
//	func() {
//		num = 5
//	}()
//
//	time.Sleep(time.Second)
//	fmt.Println(num)
//	return num
//}
//
//func Example2() func() {
//	defer fmt.Println("Deferred statement")
//
//	num := 10
//
//	a := func() {
//		num = 5
//	}
//
//	num++
//	fmt.Println(num)
//
//	return a
//}

func Example3() {

	num = 333
	defer func() {
		num++
	}()

	a := 1
	funca := func() {
		a++
	}
	funca()

	ch := make(chan struct{}, 1)

	b := make([]int, 0)
	for k, _ := range b {
		k++
	}

	ch <- struct{}{}
	return
}
