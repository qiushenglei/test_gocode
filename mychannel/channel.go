package mychannel

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func getChanData(channel chan int) {
	a := <-channel
	close(channel)
	fmt.Println(a)
}

func f1(channel chan int) {
	fmt.Println("f1进入等待")
	time.Sleep(2 * time.Second)
	channel <- 19
	// 没有close，testChan就会deadlock，因为他会一直等待其他协程往chan push元素，除非这个拿到chan的协程自旋。
	// 调用close，会通知其他监听chan的协程，读取到一个空数据
	close(channel)
}

func f2(channel chan int) {
	fmt.Println("f2进入等待")
	select {
	case val := <-channel:
		fmt.Println("f2拿到val是", val)
	}
}

func TestChan() {

	channel := make(chan int)

	go f1(channel)
	go f2(channel)
	time.Sleep(3 * time.Second)
	//fmt.Println("main go val", <-channel)
	for element := range channel {
		fmt.Println("main go  for val:", element)
	}
	return
}

// 缓冲管道发送方阻塞
func SendBlock() {
	c := make(chan int, 5)
	defer close(c)
	for i := 0; i < 4; i++ {
		//index := i
		// 如果i不使用传参方式，而是使用闭包函数，那么就会发生数据逃逸，i会被存到堆中，栈帧上的i变成指针指向堆，导致协程里的i不一定打印0,1,2,3
		go func(i int) {
			c <- i
			fmt.Printf("i=%d成功插入chan\n", i)
		}(i)
	}
	time.Sleep(3 * time.Second)

	//for v := range c {
	//	fmt.Println(v)
	//}

	for j := 0; j < 4; j++ {
		fmt.Println(<-c)
	}
}

// 缓冲管道发送方阻塞
func SendBlock1() {
	c := make(chan int, 3)
	defer close(c)
	for i := 0; i < 4; i++ {
		// 如果i不使用传参方式，而是使用闭包函数，那么就会发生数据逃逸，i会被存到堆中，栈帧上的i变成指针指向堆，导致协程里的i不一定打印0,1,2,3
		go func(i int) {
			c <- i
			fmt.Printf("i=%d成功插入chan\n", i)
		}(i)
	}
	time.Sleep(3 * time.Second)

	//打印，2协程阻塞等待
	//i=3成功插入chan
	//i=0成功插入chan
	//i=1成功插入chan

}

func SendBlock2() {
	c := make(chan int)
	go func() {
		a := <-c
		fmt.Println(a)

		a = <-c
		fmt.Println(a)
	}()
	fmt.Println("sleep 3 sec")
	time.Sleep(3 * time.Second)

	select {
	//hchan.recvq没有等待的g，
	// 如果有下一个case：虽然判定为阻塞，但是此g不会被加到sendq，而是跳到下一个case。
	// 如果下一个case也被阻塞并且还没有default，看谁哪个case先被唤醒
	// 如果没有下一个case：则阻塞，此g被加到sendq
	case c <- 2:
		fmt.Println("case1")
	// 下一个
	case c <- 3:
		fmt.Println("case2")
	}

	fmt.Println("sleep 5 sec")
	time.Sleep(5 * time.Second)
}

func SendBlock22() {
	c := make(chan int, 33)
	go func() {

		select {
		//hchan.recvq没有等待的g，
		// 如果有下一个case：虽然判定为阻塞，但是此g不会被加到sendq，而是跳到下一个case。
		// 如果下一个case也被阻塞并且还没有default，看谁哪个case先被唤醒
		// 如果没有下一个case：则阻塞，此g被加到sendq
		case c <- 2:
			fmt.Println("case1")
		// 下一个
		case c <- 3:
			fmt.Println("case2")
		}

	}()

	time.Sleep(3 * time.Second)
	a := <-c
	fmt.Println(a)

	a = <-c
	fmt.Println(a)

}

func SendBlock3() {
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		a := <-c
		fmt.Println(a)

		a = <-c
		fmt.Println(a)
	}()
	select {
	//hchan.recvq没有等待的g，
	// 如果有下一个case：虽然判定为阻塞，但是此g不会被加到sendq，而是跳到下一个case。
	// 如果下一个case也被阻塞并且还没有default，看谁哪个case先被唤醒
	// 如果没有下一个case：则阻塞，此g被加到sendq
	case c <- 2:
		fmt.Println("case1")
	// 下一个
	case c <- 3:
		fmt.Println("case2")
	}
	time.Sleep(time.Second)

}

type Animal struct {
	Voice  string
	nextch chan struct{}
	selfch chan struct{}
	time   int
}

func (a *Animal) Eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for a.time > 0 {
		<-a.selfch
		fmt.Println(a.Voice)
		a.time--
		if a.time <= 0 {
			close(a.nextch)
			break
		}
		a.nextch <- struct{}{}
	}
	fmt.Println(a.Voice, "end")
}

func Sequence() {

	catch := make(chan struct{})
	dogch := make(chan struct{})
	mousech := make(chan struct{})

	dog := Animal{
		"wang",
		mousech,
		dogch,
		5,
	}
	cat := Animal{
		"mi",
		dogch,
		catch,
		5,
	}
	mouse := Animal{
		"jiji",
		catch,
		mousech,
		5,
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go dog.Eat(wg)
	go cat.Eat(wg)
	go mouse.Eat(wg)

	fmt.Println("start")
	dogch <- struct{}{}
	wg.Wait()
	fmt.Println("end")
}

func IsClosed() {
	ch := make(chan struct{})

	go func() {
		close(ch)
	}()
	<-ch

	v, ok := <-ch
	if ok {
		fmt.Println("is ok", v)
	} else {
		fmt.Println("is closed", v)
	}

}

func RangeChan() {
	ch := make(chan struct{})

	go func() {
		ch <- struct{}{}
		time.Sleep(time.Second)
		fmt.Println("begin close chan")
		close(ch)
		fmt.Println("chan close done")
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

func MSender1Receiver() {
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// 随机种子，最好都加上
	rand.Seed(time.Now().Unix())

	// sender
	for i := 0; i < 2; i++ {
		index := i
		go func() {
			istr := strconv.Itoa(index)
			for {

				select {
				case <-stopCh:
					fmt.Println(istr, " num get close stopch")
					return
				case dataCh <- rand.Intn(100): //随机数
					time.Sleep(1 * time.Second)
				}

			}
		}()
	}

	// 读协程监听信号。向stopch发送close，退出sender，gc自动回收dataChan通道
	for {
		select {
		case val := <-dataCh:
			fmt.Println(val)
		case <-signalChan:
			fmt.Println("receiver get notify signal")
			close(stopCh)
			time.Sleep(3 * time.Second)
			return
		}
	}

}

func MSenderNReceiver() {
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// 随机种子，最好都加上
	rand.Seed(time.Now().Unix())

	// sender
	for i := 0; i < 2; i++ {
		index := i
		go func() {
			istr := strconv.Itoa(index)
			for {

				select {
				case <-stopCh:
					fmt.Println(istr, " num get close stopch")
					return
				case dataCh <- rand.Intn(100): //随机数
					time.Sleep(1 * time.Second)
				}

			}
		}()
	}

	// 读协程监听信号。向stopch发送close，退出sender，gc自动回收dataChan通道
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		index := i
		go func() {
			defer wg.Done()
			stri := strconv.Itoa(index)
			for {
				select {
				case val := <-dataCh:
					fmt.Println(val)
				case <-signalChan:
					fmt.Println(stri, " num receiver get notify signal, ")
					close(stopCh)
					time.Sleep(3 * time.Second)
				case <-stopCh:
					fmt.Println(stri, " num receiver start stop")
					return
				}
			}
		}()

	}
	wg.Wait()
}

func CaseInIO() {
	c := context.Background()
	c, _ = context.WithTimeout(c, time.Second*3)

	//urlA := "http://example.com/apiA"
	//urlB := "http://example.com/apiB"

	errCh := make(chan error, 1) // 创建一个用于传递错误信息的 channel

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// 发起 A 请求
		defer wg.Done()

		fmt.Println("start A request")
	}()

	go func() {
		defer wg.Done()
		// 监听 context 的 Done 信号
		select {
		case <-c.Done():
			fmt.Println("B request canceled")
		default:
			// 发起 B 请求
			fmt.Println("start B default, soon sleep 5s")
			time.Sleep(time.Second * 5)
			fmt.Println("这里会会执行吗？")
		}
	}()

	wg.Wait() // 等待两个协程完成

	// 等待错误信息或者全部请求完成
	select {
	case err := <-errCh:
		fmt.Println("Error:", err)
	default:
		fmt.Println("All requests completed successfully")
	}
}
