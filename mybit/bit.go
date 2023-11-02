package mybit

import (
	"fmt"
	"strconv"
)

func BitCal() {
	a := 4
	b := 1
	c := a | b
	fmt.Println(c)
}

func GetBin() {
	char := '我'
	codepoint := uint64(char)
	binary := strconv.FormatUint(codepoint, 2)
	fmt.Println(binary)

	a := "abcdefg"

	b := []byte(a)
	b[1] = 'h'
	fmt.Println(a, b)
	fmt.Printf("%p %p\n", &a, b)

	c := [...]int{1, 2, 3}
	d := c[:]
	fmt.Println(c, d)
	fmt.Printf("%p %p\n", &c, d)

}
