package agent

import "ubiq-cd/internal/interface-adapter/interface/connectrpc"

func Run(s connectrpc.Server) error {
	// TODO: validate options
	// TODO: Start runner and gitops daemon
	return connectrpc.Serve(s)
}
