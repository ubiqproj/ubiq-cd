package get

import (
	"context"
	"fmt"
	"io"
	"ubiq-cd/internal/interface-adapter/client"

	"github.com/spf13/cobra"
)

func NewCmdGet(client client.Client, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use: "get",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runGet(client, out)
		},
	}
}

func runGet(client client.Client, out io.Writer) error {
	res, err := client.Greet(context.Background(), "Ubiq")
	if err != nil {
		return err
	}
	fmt.Fprint(out, res)
	return nil
}
