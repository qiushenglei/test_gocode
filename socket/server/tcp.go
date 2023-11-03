package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testproj/socket/proto"
	"time"
)

type Server struct {
	parser proto.QProto
}

func (s *Server) Start() {
	// StartServer()
}

func (s *Server) Stop() {
}

func StartServer() {
	dsn := "127.0.0.1:9797"
	lis, err := net.Listen("tcp", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if liserr := lis.Close(); liserr != nil {
			fmt.Println("close liserr:", liserr)
		}
	}()

	fmt.Println("开始监听", dsn)

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func() {
		if sConnErr := conn.Close(); sConnErr != nil {
			fmt.Println("sConnErr关闭失败", sConnErr)
		}
	}()
	p := proto.NewQProto(conn)
	if err := p.Auth(); err != nil {
		fmt.Println("登录失败")
		p.Send(400, err)
		return
	}

	fmt.Println("auth success")

	// client那边还在阻塞recv，所以要回个消息过去
	if err := p.Send(200, "login success"); err != nil {
		fmt.Println("发送回包失败", err)
		return
	}

	wg := &sync.WaitGroup{}

	isStop := make(chan int)

	wg.Add(1)
	go func() {
		// 读
		defer wg.Done()
		for {
			response, err := p.Recv()
			if err != nil {
				p.Send(500, err.Error())
				isStop <- 1
				return
			}
			fmt.Println("接收到数据:", string(response.Body))
		}
	}()

	wg.Add(1)
	go func() {
		// 写
		defer wg.Done()
		tc := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-isStop:
				fmt.Println("写逻辑正常关闭")
				tc.Stop()
				return
			case <-tc.C:
				p.Send(200, "hello world")
			}
		}
	}()

	wg.Wait()

	fmt.Println("当前连接完毕")
}
