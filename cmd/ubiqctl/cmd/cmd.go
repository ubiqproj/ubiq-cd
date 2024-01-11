package cmd

import (
	"os"
	"ubiq-cd/cmd/ubiqctl/cmd/get"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "ubiqctl",
}

func init() {
	cmd.AddCommand(newCmdVersion())
	cmd.AddCommand(get.NewCmdGet())
}

func Execute() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
