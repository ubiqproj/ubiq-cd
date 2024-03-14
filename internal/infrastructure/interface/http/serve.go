package http

import (
	"net/http"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(addr string) connectrpc.Server {
	return &server{addr}
}

type server struct {
	addr string
}

func (s *server) Serve(handlers map[string]http.Handler) error {
	mux := http.NewServeMux()
	for path, handler := range handlers {
		mux.Handle(path, handler)
	}

	return http.ListenAndServe(s.addr, h2c.NewHandler(mux, &http2.Server{}))
}
