package main

import (
	"testproj/myproto"
)

func main() {
	myproto.Marshal()
	myproto.Base646Decode()
	myproto.GetEsMod()
}
