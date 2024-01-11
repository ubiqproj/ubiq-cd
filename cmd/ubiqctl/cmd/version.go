package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.0.0"

func newCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the client version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(Version)
		},
	}
}
