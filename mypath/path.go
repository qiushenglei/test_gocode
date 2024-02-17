package mypath

import (
	"fmt"
	"path"
)

func JoinPath() {
	url := "https://golang.google.cn/pkg/path"
	filename := "/aa.png"
	a := path.Join(url, filename)
	fmt.Println(a)
}
