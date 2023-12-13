package mynum

import (
	"fmt"
)

func Change() {
	a := 2
	b := 3
	c := float32(a) / float32(b)
	fmt.Printf("%T %.8f", c, c)
	//math.Ceil()
}
