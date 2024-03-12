package cmd

import (
	"io"
	"net/http"
	"ubiq-cd/cmd/ubiqctl/cmd/get"
	"ubiq-cd/internal/infrastructure/webapi/connect/gen/api/v1/apiv1connect"

	"github.com/spf13/cobra"
)

func NewUbuqctlCommand(out io.Writer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:          "ubiqctl",
		SilenceUsage: true,
	}

	cmd.AddCommand(newCmdVersion(out))

	client := apiv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	cmd.AddCommand(get.NewCmdGet(client, out))

	return cmd
}
