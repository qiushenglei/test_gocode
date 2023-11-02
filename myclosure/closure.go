package myclosure

import (
	"fmt"
	"time"
)

func DataEscape() {
	c := make(chan int)
	a := 11
	fmt.Println("main before a =", a)
	go func() {
		fmt.Println()
		time.Sleep(2 * time.Second)
		a++
		fmt.Println(a)
		close(c)
	}()
	time.Sleep(1 * time.Second)
	a++
	fmt.Println("main after a =", a)
	<-c

}

func test() {
	a := 1
	func() {
		a++
	}()
	a++
}

func Test1() func() {
	a := 1
	b := func() {
		a++
	}
	a++
	return b
}
