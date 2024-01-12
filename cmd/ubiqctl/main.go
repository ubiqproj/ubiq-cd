package main

import (
	"log/slog"
	"os"
	"ubiq-cd/cmd/ubiqctl/cmd"
)

func main() {
	cmd := cmd.NewUbuqctlCommand(os.Stdout)
	err := cmd.Execute()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
