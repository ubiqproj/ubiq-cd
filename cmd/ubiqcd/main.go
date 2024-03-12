package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	apiv1 "ubiq-cd/internal/infrastructure/webapi/connect/gen/api/v1"
	"ubiq-cd/internal/infrastructure/webapi/connect/gen/api/v1/apiv1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	slog.Error(serve().Error())
}

func serve() error {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := apiv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	return http.ListenAndServe("127.0.0.1:8080", h2c.NewHandler(mux, &http2.Server{}))
}

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[apiv1.GreetRequest],
) (*connect.Response[apiv1.GreetResponse], error) {
	slog.Info(fmt.Sprintf("Request headers: %v", req.Header()))
	res := connect.NewResponse(&apiv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Api-Version", "v1")
	return res, nil
}
