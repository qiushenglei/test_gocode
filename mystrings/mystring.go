package mystrings

import (
	"fmt"
	"strings"
)

func StringRead() {
	str := "我是string"
	fmt.Println(len(str))
	sr := strings.NewReader(str)

	// string
	var err error
	for err == nil {
		var ch rune
		var size int
		if ch, size, err = sr.ReadRune(); err == nil {
			fmt.Println(ch, size)
		}
	}

	p := make([]byte, len(str))
	if size, err := sr.Read(p); err == nil {
		fmt.Println(size)
	} else {
		fmt.Println(err.Error())
	}

	s1 := str[0:6]
	fmt.Println(s1[0:6])

	rs := []rune(str)
	fmt.Println(string(rs[0:2]))

	//bufio.NewReader()
	//r := bufio.NewReader()
	//r.Read()

}

func SubStrMbLen() {
	str := "我是发😒string"
	start, end := 0, 0

	sl := len(str)

	runeStr := []rune(str)
	rl := len(runeStr)

	unicodeCount := (sl - rl) / 2

	res := runeStr[start:end]

	fmt.Println(string(res), unicodeCount)
}

func SubStrByByteLen() {
	str := "我是发😒string"
	start, end := 0, 0

	sl := len(str)

	runeStr := []rune(str)
	rl := len(runeStr)

	unicodeCount := (sl - rl) / 2

	res := runeStr[start:end]

	fmt.Println(string(res), unicodeCount)
}

func Postion() {
	str := "我是string是😘"
	pos := strings.Index(str, "是s")
	fmt.Println(pos)
}
