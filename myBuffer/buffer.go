package myBuffer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func BytesBuff() {
	str := "我是string，吧😍"
	byt := bytes.NewBufferString(str)
	for {
		res := byt.Next(3)
		if len(res) == 0 {
			break
		}
		fmt.Println(res)
	}
	p := make([]byte, len(str))
	byt.ReadFrom(strings.NewReader(str))
	byt.Read(p)

}

func FileBuff() {
	fi, err := os.Stat("cmd/testpkg/main.go")
	if err != nil {
		return
	}
	fp := make([]byte, fi.Size())
	file, _ := os.OpenFile("cmd/testpkg/main.go", os.O_CREATE|os.O_APPEND, 0644)
	file.Read(fp)

	//str := "我是string，吧😍"
	//file.ReadFrom(strings.NewReader(str))
}

func Bufio() {
	// bytes 和 file都实现了。这个估计是给其他未实现的使用的
	str := "nice，古柏"
	r := bufio.NewReader(strings.NewReader(str))
	p := make([]byte, len(str))
	r.Read(p)

	//w := bufio.NewWriter(strings)
	//w.ReadFrom()
}

func BufioReadSliceBlock() ([]byte, error) {
	str := "淦真的会阻塞吗？ 有没有搞错"
	b := bufio.NewReader(strings.NewReader(str))
	var line []byte
	for {
		l, more, err := b.ReadLine()
		if err != nil {
			return nil, err
		}
		// Avoid the copy if the first call produced a full line.
		if line == nil && !more {
			fmt.Println(string(l))
			return l, nil
		}
		line = append(line, l...)
		if !more {
			break
		}
	}
	return nil, nil
}
