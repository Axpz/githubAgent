package server

import (
	"fmt"
	"log"
	"sync"

	"time"

	pb "githubagent/proto/listwatcher"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"
)

const (
	Buff_Size = 10
)

type ListWatchService struct {
	pb.UnimplementedListWatchServiceServer
	mu      sync.Mutex
	clients map[string]chan *pb.Event
}

func NewListWatchServiceServer() *ListWatchService {
	return &ListWatchService{
		clients: make(map[string]chan *pb.Event),
	}
}

func (s *ListWatchService) ListWatch(stream pb.ListWatchService_ListWatchServer) error {
	req, err := stream.Recv()
	if err != nil {
		klog.Errorf("receiving message error : %v", err)
		return err
	}

	if req.Type == "register" {
		klog.Errorf("the first request type should be register, not %s", req.Type)
		return fmt.Errorf("request type error")
	}

	clientID := req.Id
	eventChan := make(chan *pb.Event, Buff_Size)
	s.mu.Lock()
	if _, ok := s.clients[clientID]; ok {
		s.mu.Unlock()
		return fmt.Errorf("duplicated register error")
	}
	s.clients[clientID] = eventChan
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, clientID)
		close(eventChan)
		s.mu.Unlock()
		klog.Infof("client device %s disconnected", clientID)
	}()

	klog.Infof("client device %s connected", clientID)

	// 模拟发送 1 个事件
	for i := 0; i < 9; i++ {
		event := &pb.Event{
			Id:        fmt.Sprintf("item-%d", i),
			Type:      "added",
			Timestamp: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
			Data:      &anypb.Any{Value: []byte(fmt.Sprintf("data-%d", i))}, // 模拟事件数据
		}

		eventChan <- event
	}

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case event := <-eventChan:
			// send to client
			if err := stream.Send(event); err != nil {
				msgErr := fmt.Errorf("failed to send event to client %s: %v", clientID, err)
				klog.Error(msgErr)
				return msgErr
			}
		default:
			// receive from clients
			req, err := stream.Recv()
			if err != nil {
				msgErr := fmt.Errorf("client %s disconnected with error: %v", clientID, err)
				klog.Error(msgErr)
				return msgErr
			}
			log.Printf("received message from client %s %s at %v type %s", clientID, req.Id, req.Timestamp, req.Type)
		}
	}
}
