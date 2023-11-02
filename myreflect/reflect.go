package myreflect

import (
	"fmt"
	"reflect"
)

type Ref struct {
}

func (f *Ref) Test0() {
	fmt.Println("is method test0")
}

func (f *Ref) Test1() {
	fmt.Println("is method test1")
}

func (f *Ref) Test2(a int) {
	fmt.Println("is method test2")
}

func IntRef() {
	s := 1235646547868769
	o := reflect.TypeOf(s)
	fmt.Println(o.Size())
	fmt.Println(o.Name())
	fmt.Println(o.NumMethod())
	fmt.Println(o.Method(0))
}

func StrRef() {
	var s *Ref
	o := reflect.TypeOf(s)
	fmt.Println(o.Size())
	fmt.Println(o.Name())
	fmt.Println(o.NumMethod())
	fmt.Println(o.Method(1))

	m, ok := o.MethodByName("Test2")
	if ok == true {
		fmt.Println(m.Index)
		fmt.Println(m.Name)
	}

	v := reflect.ValueOf(s)
	vm := v.MethodByName("Test2")
	if vm.IsValid() == true && vm.Kind() == reflect.Func {
		param := []reflect.Value{
			reflect.ValueOf(12),
		}

		vm.Call(param)
	}

	a := 'a'
	fmt.Println(a)
	fmt.Printf("%T", a)
	b := a - 26
	fmt.Printf("%T", b)
	fmt.Printf("%T", byte(b))
	aa := []byte{
		byte(a), byte(b),
	}
	fmt.Println(string(aa))
}

type S struct{}

// 该结构体类型的方法
func (s *S) MethodName() {
	fmt.Println("Method called")
}

func CallMethodToStructName() {

	s := "S"                        // 字符串
	t := reflect.TypeOf(s)          // 获取类型
	if t.Kind() == reflect.String { // 判断是否为字符串类型
		a := reflect.ValueOf(S{}).Type().String()
		fmt.Println(a)
		v := reflect.New(reflect.ValueOf(S{}).Type()).Interface()
		name := reflect.TypeOf(v).Name()
		if name == s { // 判断类型名称是否匹配
			if m := reflect.ValueOf(v).MethodByName("MethodName"); m.IsValid() { // 判断方法是否存在
				m.Call(nil) // 调用方法
			}
		}
	}
}
