package mypanic

import (
	"context"
	"fmt"
	"time"
)

func GolRecover() {
	defer func() {
		a := recover()
		fmt.Println(a, 123)

	}()

	go func() {
		panic("子协程panic")
	}()
	time.Sleep(10 * time.Second)
}

// Go 封装的go程，不用每次都手写panic
func Go(ctx context.Context, handle func(context.Context, ...interface{}), rh func(r interface{})) {
	p := func() {
		if r := recover(); r != nil {
			if rh == nil {
				return
			}
			Go(ctx, func(context.Context, ...interface{}) {
				rh(r)
			}, nil)
		}
	}

	go func() {
		defer p()
		handle(ctx)
	}()
}

func PrintValues(values ...interface{}) {
	if len(values) == 0 {
		fmt.Println("No values provided")
		return
	}

	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	PrintValues() // 传递0个参数
}
