package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/abolinjast/rover/internal/config"
)

// InitBackends initializes the backends for a server
func InitBackends(server *config.Server) error {
	for i := range server.Backends {
		backendURL, err := url.Parse(server.Backends[i].URL)
		if err != nil {
			return err
		}
		proxy := httputil.NewSingleHostReverseProxy(backendURL)
		server.Backends[i].SetProxy(proxy)
		server.Backends[i].SetAlive(true)
	}
	return nil
}
