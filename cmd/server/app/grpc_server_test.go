package app

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"githubagent/proto/listwatcher" // 导入生成的 listwatcher gRPC 包
)

func TestGrpcServer_Start_Stop(t *testing.T) {
	// 设置 gRPC 服务器配置
	opts := &GrpcServerOptions{Port: "50051"}
	grpcServer := NewGrpcServer(opts)

	// 启动 gRPC 服务器
	go func() {
		if err := grpcServer.Start(); err != nil {
			fmt.Printf("failed to start grpc server: %v", err)
		}
	}()

	// 等待一段时间，确保服务器已启动
	time.Sleep(1 * time.Second)

	// 验证服务器是否在运行
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", opts.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 进行一次简单的 gRPC 请求测试
	client := listwatcher.NewListWatchServiceClient(conn)
	_, err = client.ListWatch(context.Background())
	assert.NoError(t, err, "ListWatch call should succeed")

	// 停止服务器
	grpcServer.Stop()
}

func TestGrpcClient_ListWatch(t *testing.T) {
	// 设置 gRPC 服务器配置
	opts := &GrpcServerOptions{
		Port:     "50051",
		Services: []listwatcher.ListWatchServiceServer{&TestListWatchService{}},
	} // Set a default port if not provided
	grpcServer := NewGrpcServer(opts)

	// 启动 gRPC 服务器
	go func() {
		if err := grpcServer.Start(); err != nil {
			fmt.Printf("failed to start grpc server: %v", err)
		}
	}()

	// 等待一段时间，确保服务器已启动
	time.Sleep(1 * time.Second)

	// 进行 gRPC 请求测试
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", opts.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 模拟一个客户端，执行 ListWatch 操作
	client := listwatcher.NewListWatchServiceClient(conn)
	stream, err := client.ListWatch(context.Background())

	assert.NoError(t, err, "ListWatch call should succeed")

	// 遍历流中的所有事件
	for {
		// 读取一个事件
		event, err := stream.Recv()
		// 如果已关闭流（无更多事件），则退出
		if err == io.EOF {
			break
		}
		assert.NoError(t, err, "Failed to receive event")

		// 检查事件数据
		assert.NotNil(t, event, "Event should not be nil")
		assert.Equal(t, "added", event.Type, "Event type should be 'added'") // event.Type 根据 proto 中定义的字段
		// 可以根据需要进一步检查其他字段
		assert.NotNil(t, event.Data, "Event data should not be nil")

		fmt.Printf("timestamp:%+v", event.Timestamp)
		fmt.Printf("Data:%+v", event.Data)

		break
	}
	grpcServer.Stop()
}
