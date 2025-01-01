package app

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"githubagent/proto/listwatcher"
)

type GrpcServerOptions struct {
	Port     string
	Services []listwatcher.ListWatchServiceServer
}

// GrpcServer represents the gRPC server.
type GrpcServer struct {
	server *grpc.Server
	opts   *GrpcServerOptions
}

// NewGrpcServer creates a new instance of GrpcServer.
func NewGrpcServer(opts *GrpcServerOptions) *GrpcServer {
	s := grpc.NewServer()
	for _, service := range opts.Services {
		listwatcher.RegisterListWatchServiceServer(s, service)
	}

	return &GrpcServer{
		server: s,
		opts:   opts,
	}
}

// Start begins serving the gRPC server on the specified port.
func (s *GrpcServer) Start() error {
	address := fmt.Sprintf(":%s", s.opts.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", address, err)
	}
	log.Printf("gRPC server listening on %s", address)

	// Serve the gRPC server
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve gRPC server on %s: %v", address, err)
	}

	return nil
}

// Stop gracefully shuts down the gRPC server.
func (s *GrpcServer) Stop() {
	if s.server != nil {
		log.Println("Stopping gRPC server...")
		s.server.GracefulStop()
	} else {
		log.Println("gRPC server is already stopped or not initialized.")
	}
}

// GetServer returns the underlying gRPC server instance.
func (s *GrpcServer) GetServer() *grpc.Server {
	return s.server
}
