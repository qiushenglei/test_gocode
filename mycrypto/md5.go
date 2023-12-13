package mycrypto

import (
	"crypto/md5"
	"fmt"
)

func Md5Sum() {
	str := "abcdef"
	res := md5.Sum([]byte(str))
	fmt.Printf("%x\n", res)
}

func Md5() {
	str := "abcdef"
	h := md5.New()
	h.Write([]byte(str))
	res := h.Sum(nil)
	fmt.Printf("%x\n", res)
	//fmt.Printf(string(res))
}
