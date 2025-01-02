package app

import (
	"context"
	"fmt"
	"githubagent/internal/server"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const (
	GCPeriod = 10 * time.Minute
)

type AgentServer struct {
	HostName string
	Port     string
	NodeName string
	NodeIP   string
	GCPeriod time.Duration
}

func (s *AgentServer) StartGarbageCollection() {
	go wait.Until(func() {
		// ctx := context.Background()
		klog.InfoS("Container garbage collection succeeded")
	}, GCPeriod, wait.NeverStop)
}

func (s *AgentServer) ListenAndServe(ctx context.Context) error {

	router := server.RegisterRoutes()

	// Create an HTTP server to host the Gin router
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Use a Goroutine to start the server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// Log the error if the server fails
			fmt.Printf("Server failed: %v\n", err)
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()

	// Shutdown the server gracefully when the context is cancelled
	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
		return err
	}

	fmt.Println("Server exited gracefully.")
	return nil
}
