package greeter

import (
	"context"
	"fmt"

	apiv1 "ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1"

	"connectrpc.com/connect"
)

type Handler struct{}

func (s *Handler) Greet(
	ctx context.Context,
	req *connect.Request[apiv1.GreetRequest],
) (*connect.Response[apiv1.GreetResponse], error) {
	res := connect.NewResponse(&apiv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Api-Version", "v1")
	return res, nil
}
