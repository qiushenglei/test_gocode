package mydatatype

import (
	"fmt"
	"strings"
)

func Rune() {
	str := "邱升磊"
	fmt.Println(len(str))
	a := []byte(str)
	fmt.Println(a)

	b := []rune{'邱', '升', '磊'}
	fmt.Println(b)

	for _, v := range str {
		fmt.Println(v)
	}
}

func Split(s, seq string) []string {
	res := make([]string, 0, strings.Count(s, seq))
	start := 0
	idx := strings.Index(s, seq)
	for idx != -1 {
		res = append(res, s[start:idx])
		s = s[idx+len(seq):]
		idx = strings.Index(s, seq)
	}
	if len(s) > 0 {
		res = append(res, s)
	}
	return res
}
