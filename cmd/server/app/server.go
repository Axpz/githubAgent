package app

import (
	"context"
	"fmt"
	"githubagent/proto/listwatcher"
	"os"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/noop"

	"k8s.io/klog/v2"

	genericapiserver "k8s.io/apiserver/pkg/server"

	"githubagent/internal/server"
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

			// set up signal context for kubelet shutdown
			ctx := genericapiserver.SetupSignalContext()
			serverOption := AgentServer{}
			return Run(ctx, &serverOption)
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

func Run(ctx context.Context, s *AgentServer) error {
	klog.InfoS("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))

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

	if err := s.ListenAndServe(ctx); err != nil && ctx.Err() == nil {
		klog.ErrorS(err, "Failed to run AgentServer")
	}

	return nil
}
