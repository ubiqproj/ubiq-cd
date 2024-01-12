package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const Version = "v0.0.0"

func newCmdVersion(out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the client version",
		Run: func(_ *cobra.Command, _ []string) {
			RunVersion(out)
		},
	}
}

func RunVersion(out io.Writer) {
	fmt.Fprint(out, Version)
}
