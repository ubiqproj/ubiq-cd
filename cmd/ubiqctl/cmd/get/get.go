package get

import (
	"context"
	"fmt"
	"net/http"
	apiv1 "ubiq-cd/third_party/connect/gen/api/v1"
	"ubiq-cd/third_party/connect/gen/api/v1/apiv1connect"

	"connectrpc.com/connect"
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	return &cobra.Command{
		Use:  "get",
		RunE: get,
	}
}

func get(_ *cobra.Command, _ []string) error {
	client := apiv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&apiv1.GreetRequest{Name: "Ubiq"}),
	)
	if err != nil {
		return err
	}
	fmt.Println(res.Msg.Greeting)
	return nil
}
