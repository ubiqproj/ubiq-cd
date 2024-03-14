package cmd

import (
	"io"
	"net/http"
	"ubiq-cd/internal/infrastructure/cmd/ubiqctl/cmd/get"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1/apiv1connect"

	"github.com/spf13/cobra"
)

func NewCmdUbiqctl(out io.Writer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "ubiqctl",
		SilenceUsage: true,
	}

	client, err := newClientByUrl(cmd)
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(newCmdVersion(out))
	cmd.AddCommand(get.NewCmdGet(client, out))

	return cmd, nil
}

func newClientByUrl(cmd *cobra.Command) (apiv1connect.GreetServiceClient, error) {
	err := cmd.MarkFlagRequired("url")
	if err != nil {
		return nil, err
	}
	url := cmd.Flags().String("url", "http://localhost:8080", "The URL for the UbiqCD agent.")
	// TODO: validate flag value
	client := apiv1connect.NewGreetServiceClient(
		http.DefaultClient,
		*url,
	)
	return client, nil
}
