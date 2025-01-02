package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func mainTest() {
	// 父级上下文，模拟控制整个操作生命周期
	parentCtx, cancelParent := context.WithCancel(context.Background())

	// 第二层，子级上下文：与父级上下文关联，作为子任务的生命周期控制
	childCtx, _ := context.WithCancel(parentCtx)

	// 第三层，孙级上下文：与子级上下文关联，作为孙任务的生命周期控制
	grandChildCtx, _ := context.WithCancel(childCtx)

	// 启动父级任务
	go func() {
		fmt.Println("Parent task started")
		select {
		case <-parentCtx.Done():
			fmt.Println("Parent task canceled")
		}
	}()

	// 启动子级任务
	go func() {
		fmt.Println("Child task started")
		select {
		case <-childCtx.Done():
			fmt.Println("Child task canceled")
		}
	}()

	// 启动孙级任务
	go func() {
		fmt.Println("Grandchild task started")
		select {
		case <-grandChildCtx.Done():
			fmt.Println("Grandchild task canceled")
		}
	}()

	// 模拟父级任务运行一段时间
	time.Sleep(2 * time.Second)

	// 取消父级任务，这将触发子级和孙级任务的取消
	cancelParent()
	// 如果想单独取消子级或孙级任务，调用它们的 cancel 函数：
	// cancelChild()
	// cancelGrandChild()

	// 程序继续等待，直到所有任务完成
	// time.Sleep(2 * time.Second)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// 等待接收到终止信号（SIGINT 或 SIGTERM）
	<-sigCh

	fmt.Println("Main function ends")
}
