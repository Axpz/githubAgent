package main

import (
	"fmt"
	"githubagent/cmd/server/app"
	"os"
)

func main() {
	cmd := app.NewCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
