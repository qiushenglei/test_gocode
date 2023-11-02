package err

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func HttpErr() {
	// 创建一个context和cancel函数
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 使用errgroup创建一个新的Group，并将context传递给它
	var g errgroup.Group
	g.Go(func() error {
		fmt.Println("Starting server...")
		http.ListenAndServe(":8080", nil)
		select {}
		fmt.Println("这里因该不会结束把")
		return nil
	})
	g.Go(func() error {
		fmt.Println("Performing task 1...")
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// 执行某些操作
			// 如果结果错误
			return fmt.Errorf("Task 1 error")
		}
	})
	g.Go(func() error {
		//time.Sleep(5 * time.Second)
		//return errors.New("5秒过了，返回err")
		return nil
	})

	// 等待所有任务完成或者任意一个任务报错
	err := g.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		// 取消其他正在运行的任务
		cancel()
	}

	fmt.Println("All tasks are done")
}

func HttpErr1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// 启动 HTTP 服务器协程
	g.Go(func() error {
		fmt.Println("Starting server...")
		server := &http.Server{Addr: ":8080"}
		go func() {
			// 监听退出信号
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			select {
			case <-ctx.Done():
				// 收到取消信号时关闭服务器
				fmt.Println("Shutting down server...")
				server.Shutdown(ctx)
			case sig := <-quit:
				// 收到退出信号时关闭服务器
				fmt.Printf("Received signal %s, shutting down...\n", sig)
				server.Shutdown(ctx)
			}
		}()

		return server.ListenAndServe()
	})

	// 子协程执行某个任务
	g.Go(func() error {
		fmt.Println("Performing task...")
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// 执行某些操作
			// 如果结果错误
			return fmt.Errorf("Task error")
		}
	})

	err := g.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("All tasks are done")
}
