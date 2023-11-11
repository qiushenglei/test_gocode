package main

import (
	"fmt"
	"reflect"
	"testproj/mychannel"
)

func main() {
	mychannel.Sequence()
	//mysync.SafeMap()
	//mysync.UnsafeMap()
	//mycrypto.Md5()
	//myinterface.Interface3()
	//myinterface.Interface2()
	//mychannel.SendBlock2()
	//mychannel.SendBlock()
	//mychannel.SendBlock3()
	//mydefer.Panic()
	//res := mydefer.Example1()
	//fmt.Println(res)
	//myinterface.Operate()
	//mymap.Cap()
	//mycontext.WithTimeout()
	//mycontext.WithCancel1()
	//mydatatype.Rune()
	//mymap.Cap()
	//mymap.Cap1()
	//mymap.Copy()
	//mymap.Delete()
	//myslice.Delete()
	//myslice.IsParam()
	//myslice.IsParam1()
	//myslice.ArrToSlice()
}

func b() {
	a := map[string]interface{}{
		"a": 1,
		"b": 12,
	}

	b := map[string]interface{}{
		"a": 1,
		"b": "asd",
	}
	if reflect.DeepEqual(a, b) {
		fmt.Println("is equal")
	} else {
		fmt.Println("not equal")
	}
}
