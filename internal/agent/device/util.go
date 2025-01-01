package device

import (
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

func retryOrForget[T comparable](logger klog.Logger, queue workqueue.TypedRateLimitingInterface[T], obj T, requeue bool, maxRetries int) {
	if !requeue {
		queue.Forget(obj)
		return
	}

	requeueCount := queue.NumRequeues(obj)
	if requeueCount < maxRetries {
		queue.AddRateLimited(obj)
		return
	}

	logger.V(4).Info("retried several times", "obj", obj, "count", requeueCount)
	queue.Forget(obj)
}
