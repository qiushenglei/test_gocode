package main

import (
	"fmt"
	"reflect"
	"testproj/mycontext"
)

func main() {
	mycontext.WithTimeout()
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
