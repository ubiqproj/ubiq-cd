package client

import (
	"context"
	apiv1 "ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1"

	"connectrpc.com/connect"
)

func (c *client) Greet(ctx context.Context, name string) (string, error) {
	res, err := c.GreetServiceClient.Greet(
		ctx,
		connect.NewRequest(&apiv1.GreetRequest{Name: "Ubiq"}),
	)
	if err != nil {
		return "", err
	}
	return res.Msg.Greeting, nil
}
