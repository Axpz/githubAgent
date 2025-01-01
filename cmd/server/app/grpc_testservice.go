package app

import (
	"fmt"
	"time"

	lw "githubagent/proto/listwatcher"

	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type TestListWatchService struct {
	lw.UnimplementedListWatchServiceServer
}

func (s *TestListWatchService) ListWatch(stream lw.ListWatchService_ListWatchServer) error {
	// 模拟发送 1 个事件
	for i := 0; i < 1; i++ {
		event := &lw.Event{
			Id:        fmt.Sprintf("item-%d", i),
			Type:      "added",
			Timestamp: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
			Data:      &anypb.Any{Value: []byte(fmt.Sprintf("data-%d", i))}, // 模拟事件数据
		}

		// 发送事件流
		if err := stream.Send(event); err != nil {
			return fmt.Errorf("failed to send event: %v", err)
		}

		// 模拟延时
		time.Sleep(1 * time.Second)
	}

	return nil
}
