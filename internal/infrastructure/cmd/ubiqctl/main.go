package ubiqctl

import (
	"os"
	"ubiq-cd/internal/infrastructure/cmd/ubiqctl/cmd"
)

func Run() error {
	cmd, err := cmd.NewCmdUbiqctl(os.Stdout)
	if err != nil {
		return err
	}
	return cmd.Execute()
}
