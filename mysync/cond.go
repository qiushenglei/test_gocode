package mysync

import (
	"fmt"
	"sync"
	"time"
)

func CondSync() {
	cond := sync.NewCond(&sync.Mutex{})
	var ready bool

	// 等待 goroutine
	go func() {
		cond.L.Lock()

		for !ready {
			fmt.Println("Waiting...")
			cond.Wait()
			fmt.Println("Woke up!")
		}
		cond.L.Unlock()
	}()

	time.Sleep(time.Second) // 模拟一些操作

	// 通知 goroutine
	cond.L.Lock()
	fmt.Println("Notify")
	ready = true
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second) // 模拟一些操作
}
