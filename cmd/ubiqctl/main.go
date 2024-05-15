package main

import (
	"log/slog"
	"os"
	"ubiq-cd/internal/infrastructure/cmd/ubiqctl"
)

func main() {
	err := ubiqctl.Run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
