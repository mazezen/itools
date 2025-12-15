package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mazezen/itools"
)

func main() {
	// 创建调度器
	scheduler := itools.NewScheduler()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 示例任务1：5秒后执行
	scheduler.AddTask(1, "Task 1", time.Now().Add(5*time.Second), func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Task 1 canceled: %v\n", ctx.Err())
			return
		default:
			fmt.Println("Task 1 is running!")
		}
	})

	// 示例任务2：10秒后执行
	scheduler.AddTask(2, "Task 2", time.Now().Add(10*time.Second), func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Task 2 canceled: %v\n", ctx.Err())
			return
		default:
			fmt.Println("Task 2 is running with complex logic!")
		}
	})

	// 启动调度器
	fmt.Println("Scheduler started...")
	go scheduler.Run(ctx)

	// 模拟取消任务
	time.Sleep(2 * time.Second)
	scheduler.CancelTask(1)

	// 等待任务完成
	time.Sleep(15 * time.Second)
	cancel() // 停止调度器
	fmt.Println("Main function exiting...")
}
