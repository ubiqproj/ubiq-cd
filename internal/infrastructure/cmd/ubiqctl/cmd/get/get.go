package get

import (
	"context"
	"fmt"
	"io"
	apiv1 "ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1/apiv1connect"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

func NewCmdGet(client apiv1connect.GreetServiceClient, out io.Writer) *cobra.Command {
	return &cobra.Command{
		Use: "get",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runGet(client, out)
		},
	}
}

func runGet(client apiv1connect.GreetServiceClient, out io.Writer) error {
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&apiv1.GreetRequest{Name: "Ubiq"}),
	)
	if err != nil {
		return err
	}
	fmt.Fprint(out, res.Msg.Greeting)
	return nil
}
