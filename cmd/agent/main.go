package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
	"k8s.io/klog/v2"

	"githubagent/internal/agent/device"
	"githubagent/internal/watch"
)

type options struct {
	flags         []cli.Flag
	configFile    string
	kubeletSocket string
}

func main() {
	c := cli.NewApp()
	o := &options{}
	c.Name = "Agent"
	c.Usage = "Agent for Kubernetes"
	c.Version = "v0.0.1"
	c.Action = func(ctx *cli.Context) error {
		return start(ctx, o)
	}

	c.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "ca-certs",
			Value:   "",
			Usage:   "the configuration of certs",
			EnvVars: []string{"CA_CERTS"},
		},

		&cli.StringFlag{
			Name:        "config-file",
			Usage:       "the path to a config file as an alternative to command line options or environment variables",
			Destination: &o.configFile,
			EnvVars:     []string{"CONFIG_FILE"},
		},
	}
	o.flags = c.Flags

	err := c.Run(os.Args)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

func start(c *cli.Context, o *options) error {
	klog.InfoS(fmt.Sprintf("Starting %s", c.App.Name), "version", c.App.Version)

	// watch the configuration file
	configDir := filepath.Dir(o.configFile)
	klog.Infof("Starting FS watcher for %v", configDir)
	watcher, err := watch.Files(configDir)
	if err != nil {
		return fmt.Errorf("failed to create FS watcher for %s: %v", configDir, err)
	}
	defer watcher.Close()

	klog.Info("Starting OS watcher.")
	sigs := watch.Signals(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var started bool
	var restartTimeout <-chan time.Time
	var devices []device.Interface
restart:
	// If we are restarting, stop devices from previous run.
	if started {
		err := stopDevices(devices)
		if err != nil {
			return fmt.Errorf("error stopping devices from previous run: %v", err)
		}
	}

	klog.Info("Starting devices.")
	devices, restartdevices, err := startDevices(c, o)
	if err != nil {
		return fmt.Errorf("error starting devices: %v", err)
	}
	started = true

	if restartdevices {
		klog.Infof("Failed to start one or more devices. Retrying in 30s...")
		restartTimeout = time.After(30 * time.Second)
	}

	// Start an infinite loop, waiting for several indicators to either log
	// some messages, trigger a restart of the devices, or exit the program.
	for {
		select {
		// If the restart timeout has expired, then restart the devices
		case <-restartTimeout:
			goto restart

		// Detect a restart by watching for a newly created file. When this occurs, restart this loop,
		// restarting all of the devices in the process.
		case event := <-watcher.Events:
			if o.configFile != "" && (event.Op&fsnotify.Create) == fsnotify.Create {
				klog.Infof("inotify: %s created, restarting.", o.configFile)
				goto restart
			}

		// Watch for any other fs errors and log them.
		case err := <-watcher.Errors:
			klog.Infof("inotify: %s", err)

		// Watch for any signals from the OS. On SIGHUP, restart this loop,
		// restarting all of the devices in the process. On all other
		// signals, exit the loop and exit the program.
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				klog.Info("Received SIGHUP, restarting.")
				goto restart
			default:
				klog.Infof("Received signal \"%v\", shutting down.", s)
				goto exit
			}
		}
	}
exit:
	err = stopDevices(devices)
	if err != nil {
		return fmt.Errorf("error stopping devices: %v", err)
	}

	// sleep a while
	time.Sleep(3 * time.Second)

	return nil
}

func startDevices(ctx *cli.Context, o *options) ([]device.Interface, bool, error) {
	// Load the configuration file
	klog.Info("Loading configuration.")
	config, err := loadConfig(ctx, o)
	if err != nil {
		return nil, false, fmt.Errorf("unable to load config: %v", err)
	}

	klog.Infof("\nRunning with config:\n%v", config)

	// Get the set of devices.
	klog.Info("Retrieving devices.")
	devices := getDevices(config)

	// Loop through all devices, starting them if they have any devices
	// to serve. If even one device fails to start properly, try
	// starting them all again.
	started := 0
	for _, d := range devices {
		if err := d.Start(ctx.Context); err != nil {
			klog.Errorf("Failed to start device: %v", err)
			return devices, true, nil
		}
		started++
	}

	if started == 0 {
		klog.Info("No devices found. Waiting indefinitely.")
	}

	return devices, false, nil
}

func stopDevices(devices []device.Interface) error {
	klog.Info("Stopping devices.")
	var errs error
	for _, p := range devices {
		errs = errors.Join(errs, p.Stop())
	}

	return errs
}

// Getdevices returns a set of devices for the specified configuration.
func getDevices(config any) []device.Interface {
	return device.GetAllDevices()
}

func loadConfig(c *cli.Context, o *options) (any, error) {
	// TODO
	return nil, nil
}
