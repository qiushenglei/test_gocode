package mymutex

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

func MutexWaiting() {

	f, err := os.Create("profile")
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer f.Close()
	defer trace.Stop()

	f1, err := os.Create("pprof")
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(f1); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	lock := sync.Mutex{}
	go func() {
		lock.Lock()
		time.Sleep(5 * time.Second)
		fmt.Println("处理业务A")
		lock.Unlock()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		lock.Lock()
		fmt.Println("处理业务B")
		lock.Unlock()
	}()
	fmt.Println("主线完成sleep")
	time.Sleep(10 * time.Second)
	fmt.Println("主线完成")
}

func MutexOrder() {

	m := sync.Mutex{}
	go user1(&m)
	go user2(&m)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	select {
	case <-signalChan:
		fmt.Println("catch interrupt signal")
		break
	}
}

func printer(str string, m *sync.Mutex) {
	m.Lock()         //加锁
	defer m.Unlock() //解锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 1)
	}
}
func user1(m *sync.Mutex) {
	printer("hello ", m)
}
func user2(m *sync.Mutex) {
	printer("world", m)
}

func RMutex() {
	ch := make(chan struct{})
	rw := &sync.RWMutex{}
	go func() {
		rw.RLock()
		fmt.Println("fun1")
		go func() {
			rw.RLock()
			fmt.Println("fun3")
			rw.RUnlock()
		}()
		rw.RUnlock()
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		rw.Lock()
		defer rw.Unlock()
		fmt.Println("fun2")
		close(ch)
	}()

	<-ch
}

func test() {
	i := 1
	go func() {
		i++
	}()

	go func() {
		i++
	}()
}
