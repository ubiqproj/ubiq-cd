package main

import (
	"log/slog"
	"os"
	"ubiq-cd/internal/infrastructure/cmd/ubiqcd"
)

func main() {
	err := ubiqcd.Run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
