package mycontext

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"sync"
	"time"
)

func WithCancel() {
	ctx1, cancel := context.WithCancel(context.TODO())
	ctx := context.WithValue(ctx1, "key", "value")
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()

		for {
			select {
			case val := <-ctx.Done():
				fmt.Println(val)
				fmt.Println("this is func1 end")
				return
			default:
				fmt.Println("this is func1")
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case val := <-ctx.Done():
				fmt.Println(val)
				fmt.Println("this is func2 end")
				return
			default:
				fmt.Println("this is func2")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()

}

func WithCancel1() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("finish")
				time.Sleep(time.Millisecond * 1)
				for i := 0; i < 10; i++ {
					fmt.Println(i)
				}
			}
		}
	}()

	time.Sleep(time.Millisecond * 100)
	cancel()
	time.Sleep(time.Millisecond * 13)
}

func WithTimeout() {
	c, _ := context.WithTimeout(context.Background(), time.Second*5)
	ch := make(chan struct{})
	go func() {
		i := 0
		for {
			select {
			case <-c.Done():
				log.Print("child goroutine is finish")
				close(ch)
				return
			default:
				i++
				fmt.Println(i)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	fmt.Println("main goroutine is blocked")
	<-ch
	//cancel()
}
