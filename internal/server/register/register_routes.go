package register

import (
	"sync"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

var (
	registry = map[string]func(*gin.Engine){}
	mu       sync.Mutex
)

func Register(name string, fn func(*gin.Engine)) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := registry[name]; ok {
		klog.Fatalf("Route %s already registered", name)
	}

	registry[name] = fn
}

func Router() *gin.Engine {
	router := gin.Default()

	mu.Lock()
	defer mu.Unlock()

	for k, fn := range registry {
		fn(router)
		klog.Infof("Registering route: %s", k)
	}

	return router
}
