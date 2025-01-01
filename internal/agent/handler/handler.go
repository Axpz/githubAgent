package handler

import (
	"context"

	"k8s.io/klog/v2"

	pb "githubagent/proto/listwatcher"
)

func Hello(ctx context.Context, v any) []byte {
	e := v.(*pb.Event)
	logger := klog.FromContext(ctx)
	logger.Info("hello %+v", v)
	return []byte("hello " + e.Id)
}
