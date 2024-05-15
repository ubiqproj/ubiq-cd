package ubiqcd

import (
	"fmt"
	"ubiq-cd/internal/infrastructure/interface/http"
	"ubiq-cd/internal/interface-adapter/agent"

	"github.com/spf13/pflag"
)

func Run() error {
	return agent.Run(http.NewServer(serverOption()))
}

func serverOption() string {
	host := pflag.String("host", "127.0.0.1", "")
	port := pflag.StringP("port", "p", "8080", "")
	pflag.Parse()
	return fmt.Sprintf("%s:%s", *host, *port)
}
