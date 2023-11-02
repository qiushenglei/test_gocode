package err

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func CancelTask() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// 子协程执行某个任务
	g.Go(func() error {
		fmt.Println("Performing task...")
		for {
			select {
			case <-ctx.Done():
				// 如果收到取消信号，则停止任务执行
				fmt.Println("Task stopped")
				return ctx.Err()
			default:
				// 模拟某些操作
				time.Sleep(10 * time.Second)
				fmt.Println("Task completed")
				return nil
			}
		}

	})

	g.Go(func() error {
		fmt.Println("Performing task 1...")

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// 执行某些操作
			// 如果结果错误
			time.Sleep(2 * time.Second)
			fmt.Println("task 2 sleep finish")
			return fmt.Errorf("Task 1 error")
		}
	})

	// 等待所有子协程退出
	err := g.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("All tasks are done")
}
