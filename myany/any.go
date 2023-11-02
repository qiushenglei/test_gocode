package myany

import "fmt"

func TAny() {
	aa := map[int]any{
		1: "a",
		2: 2,
		3: []int{1, 33},
	}

	fmt.Println(aa)
}
