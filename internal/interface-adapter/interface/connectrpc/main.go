package connectrpc

import (
	"net/http"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1/apiv1connect"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/handler/greeter"
)

type Server interface {
	Serve(handlers map[string]http.Handler) error
}

func Serve(s Server) error {
	return s.Serve(handlers())
}

func handlers() map[string]http.Handler {
	handlers := map[string]http.Handler{}
	path, handler := apiv1connect.NewGreetServiceHandler(&greeter.Handler{})
	handlers[path] = handler
	return handlers
}
