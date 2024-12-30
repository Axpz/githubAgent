package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// type subscriber
type Subscriber struct {
	id int
	ch chan string // 每个订阅者有一个独立的通道
}

type PubSub struct {
	subscribers []Subscriber
	mu          sync.Mutex
	ch          chan string
}

func (ps *PubSub) Subscribe(sub Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subscribers = append(ps.subscribers, sub)
}

func (ps *PubSub) Publish(message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// 向多个subscriber添加消息
	for _, sub := range ps.subscribers {
		// 启动一个携程向channel添加消息，其中channel是thread safe所以可以直接执行
		go func(s Subscriber) {
			s.ch <- message
		}(sub)
	}
}

// subscriber listen
func (sub *Subscriber) Listen(ctx context.Context) {
	for {
		select {
		case msg := <-sub.ch:
			fmt.Printf("subscriber %d received: %s\n", sub.id, msg)
		case <-ctx.Done():
			fmt.Printf("subscriber %d exiting\n", sub.id)
			return
		}
	}
}

func mainPubSub() {
	pubSub := PubSub{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sub1 := Subscriber{id: 1, ch: make(chan string)}
	sub2 := Subscriber{id: 2, ch: make(chan string)}
	sub3 := Subscriber{id: 3, ch: make(chan string)}

	pubSub.Subscribe(sub1)
	pubSub.Subscribe(sub2)
	pubSub.Subscribe(sub3)

	go sub1.Listen(ctx)
	go sub2.Listen(ctx)
	go sub3.Listen(ctx)

	pubSub.Publish("hello 1")
	pubSub.Publish("hello 2")

	// 使用 os/signal 捕获系统信号（如 Ctrl+C）
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// 等待接收到终止信号（SIGINT 或 SIGTERM）
	<-sigCh
	fmt.Println("Received termination signal, shutting down...")
	fmt.Print("hello end")
}
