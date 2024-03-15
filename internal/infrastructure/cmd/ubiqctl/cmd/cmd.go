package cmd

import (
	"io"
	"net/http"
	"ubiq-cd/internal/infrastructure/cmd/ubiqctl/cmd/get"
	"ubiq-cd/internal/interface-adapter/client"

	"github.com/spf13/cobra"
)

func NewCmdUbiqctl(out io.Writer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "ubiqctl",
		SilenceUsage: true,
	}

	client := newClientByUrl(cmd)

	cmd.AddCommand(newCmdVersion(out))
	cmd.AddCommand(get.NewCmdGet(client, out))

	return cmd, nil
}

func newClientByUrl(cmd *cobra.Command) client.Client {
	url := cmd.Flags().String("url", "http://localhost:8080", "The URL for the UbiqCD agent.")
	// TODO: validate flag value
	return client.New(http.DefaultClient, *url)
}
