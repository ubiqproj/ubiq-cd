package client

import (
	"context"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1/apiv1connect"

	"connectrpc.com/connect"
)

type Client interface {
	Greet(ctx context.Context, name string) (string, error)
}

type client struct {
	GreetServiceClient apiv1connect.GreetServiceClient
}

func New(httpClient connect.HTTPClient, baseUrl string) Client {
	return &client{apiv1connect.NewGreetServiceClient(
		httpClient,
		baseUrl,
	)}
}
