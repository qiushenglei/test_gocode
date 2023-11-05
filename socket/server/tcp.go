package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	fmt.Println("listening to 9797")

	ctx, cancel := context.WithCancel(context.Background())
	// 注册信号
	go ListenSignal(lis, cancel)

	for {
		conn, err := lis.Accept()
		fmt.Println("connect to ", conn.RemoteAddr().String())
		if err != nil {
			log.Println(err)
			continue
		}

		// 工作协程
		go process(ctx, conn)
	}
}

func process(ctx context.Context, conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("ConnErr关闭失败", err.Error())
		}
	}()

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("get panic")
		}
	}()

	childCtx, cancel := context.WithCancel(ctx)
	p := proto.NewQProto(childCtx, conn, cancel)
	if err := p.Auth(); err != nil {
		fmt.Println("鉴权失败")
		p.Send(400, err)
		return
	} else {
		p.Send(200, "鉴权成功")
	}

	eg, ctx := errgroup.WithContext(childCtx)

	// 双工通道
	eg.Go(Send(ctx, p))
	eg.Go(Recv(ctx, p))

	if err := eg.Wait(); err != nil {
		fmt.Println("协程嘎了", err.Error())
	}

}

func ListenSignal(listener net.Listener, cancel context.CancelFunc) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	fmt.Println("接收到关闭信号")
	cancel()

	t := time.NewTimer(5 * time.Second)
	<-t.C
	listener.Close()
}

func Send(c context.Context, p *proto.QProto) func() error {
	return func() error {
		defer fmt.Println("defer Send协程关闭")
		//监听关闭信号
		var isDone bool
		go func() {
			select {
			case <-c.Done():
				fmt.Println("send 接收到done信号")
				isDone = true
				return
			}
		}()

		for {
			if isDone {
				fmt.Println("接收到关闭信号，Send协程关闭")
				return nil
			}
			time.Sleep(3 * time.Second)
		}

		return nil
	}
}

func Recv(c context.Context, p *proto.QProto) func() error {
	return func() error {
		defer fmt.Println("defer Recv协程关闭")
		// 监听关闭信号
		var isDone bool
		go func() {
			select {
			case <-c.Done():
				fmt.Println("recv 接收到done信号")
				isDone = true
				return
			}
		}()

		for {
			if isDone {
				fmt.Println("接收到关闭信号，Recv协程关闭")
				return nil
			}

			response, err := p.Recv()
			if errors.Is(err, errorpkg.ErrSystem) {
				fmt.Println("recv失败，closed???")
				p.Close("recv 失败")
				return err
			} else if errors.Is(err, errorpkg.ErrParam) {
				p.Send(500, "参数错误")
				return nil
			}
			fmt.Println("接收到", string(response.Body))

			// 处理业务

			// 发送
			p.Send(200, string(response.Body)+"ok")
		}
	}
}
