package device

import (
	"context"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	"githubagent/internal/agent/handler"
	pb "githubagent/proto/listwatcher"
)

// Devices wraps a map[string]*Device with some functions.
type Devices map[string]*Device

var (
	devs = Devices{}
	lock sync.RWMutex
)

func init() {
	d := New(handler.Hello)
	Add(d)
}

// Device wraps pluginapi.Device with extra metadata and functions.
type Device struct {
	q          workqueue.TypedRateLimitingInterface[any]
	maxRetries int
	workers    int
	handler    func(ctx context.Context, v any) []byte
	stopCh     chan struct{}

	conn   *grpc.ClientConn
	stream grpc.BidiStreamingClient[pb.Event, pb.Event]

	ID                string
	Name              string
	Health            string
	Paths             []string
	Index             string
	TotalMemory       uint64
	ComputeCapability string
	// Replicas stores the total number of times this device is replicated.
	// If this is 0 or 1 then the device is not shared.
	Replicas int
}

// deviceInfo defines the information the required to construct a Device
type deviceInfo interface {
	GetUUID() (string, error)
	GetPaths() ([]string, error)
	GetNumaNode() (bool, int, error)
	GetTotalMemory() (uint64, error)
	GetComputeCapability() (string, error)
}

func New(handler func(ctx context.Context, v any) []byte) *Device {
	return &Device{
		q: workqueue.NewTypedRateLimitingQueue[any](
			workqueue.DefaultTypedControllerRateLimiter[any]()),
		workers: 1,
		handler: handler,
		stopCh:  make(chan struct{}),
	}
}

func (d *Device) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	go wait.UntilWithContext(ctx, func(ctx context.Context) {
		// conn, err := grpc.NewClient("dns:///localhost:50051")
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			klog.Errorf("failed to create client connection: %v", err)
			return
		}

		client := pb.NewListWatchServiceClient(conn)
		stream, err := client.ListWatch(ctx)
		if err != nil {
			conn.Close()
			klog.Errorf("failed to create stream: %v", err)
			return
		}

		d.conn = conn
		d.stream = stream

		// register
		stream.Send(&pb.Event{
			Id: d.ID,
		})

		for {
			event, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					klog.Error("Stream closed by server")
				} else if status.Code(err) == codes.Aborted {
					klog.Error("Stream interrupted: ", err)
				} else {
					klog.Errorf("failed to receive: %v", err)
				}
				return
			}

			d.q.Add(event)
		}

	}, 5*time.Second)

	workers := d.workers

	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, func(ctx context.Context) {
			v, quit := d.q.Get()
			if quit {
				return
			}
			defer d.q.Done(v)

			logger := klog.FromContext(ctx)
			// Track whether or not we should retry this sync
			retry := false
			defer func() {
				retryOrForget(logger, d.q, v, retry, d.maxRetries)
			}()

			resp := d.handler(ctx, v)
			if err := d.stream.Send(&pb.Event{
				Id:   d.ID,
				Data: &anypb.Any{Value: resp},
			}); err != nil {
				klog.Error(err)
			}

		}, 0)
	}

	go func() {
		<-d.stopCh
		klog.Infof("Shutting down %d workers", workers)
		cancel()
	}()

	klog.Infof("Start device %s with %d workers done", d.ID, workers)
	return nil
}

func (d *Device) Stop() error {
	// shutdown the queue for quit
	d.q.ShutDownWithDrain()

	// to cancel the context
	close(d.stopCh)

	d.conn.Close()

	klog.Infof("Stop device %s", d.ID)

	return nil
}
