package app

import (
	"context"
	"fmt"
	"githubagent/proto/listwatcher"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/noop"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/klog/v2"

	"githubagent/internal/server"
	_ "githubagent/internal/server/handlers"
	_ "githubagent/internal/server/handlers/github"
	"githubagent/internal/server/register"
)

func init() {
	otel.SetMeterProvider(noop.NewMeterProvider())
}

const (
	conponentName = "agentServer"
)

// NewCommand creates a *cobra.Command object with default parameters
func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:                conponentName,
		Long:               `Start the agent server`,
		DisableFlagParsing: true,
		SilenceUsage:       true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}

	// ugly, but necessary, because Cobra's default UsageFunc and HelpFunc pollute the flagset with global flags
	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine(), "null")
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine(), "null")
	})

	return cmd
}

func Run() error {
	klog.InfoS("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))

	// set up signal context for kubelet shutdown
	ctx := genericapiserver.SetupSignalContext()

	grpcOptions := GrpcServerOptions{
		Port:     "50051",
		Services: []listwatcher.ListWatchServiceServer{server.NewListWatchServiceServer()},
	}
	grpcs := NewGrpcServer(&grpcOptions)
	go func() {
		if err := grpcs.Start(); err != nil {
			fmt.Printf("failed to start grpc server: %v", err)
		}
	}()

	if err := listenAndServe(ctx); err != nil && ctx.Err() == nil {
		klog.ErrorS(err, "Failed to run AgentServer")
	}

	return nil
}

func listenAndServe(ctx context.Context) error {

	router := register.Router()

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
