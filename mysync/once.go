package mysync

import (
	"fmt"
	"sync"
	"time"
)

func Once() {
	o := sync.Once{}
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("sleep done")
		o.Do(func() {
			fmt.Println("is go 1")
			//<-ch
		})
		fmt.Println("go1 done")
	}()

	go func() {
		defer wg.Done()
		o.Do(func() {
			fmt.Println("is go 2")
			time.Sleep(time.Second * 3)
			ch <- struct{}{}
		})
		fmt.Println("go2 done")
	}()
	wg.Wait()
}
