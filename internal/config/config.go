package config

import (
	"net/http/httputil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Backend struct {
	URL   string `yaml:"url"`
	proxy *httputil.ReverseProxy
	alive bool
	mux   sync.RWMutex
}

type Server struct {
	Name     string    `yaml:"name"`
	Port     int       `yaml:"port"`
	Backends []Backend `yaml:"backends"`
	index    int
	mux      sync.Mutex
}

type Config struct {
	Servers []Server `yaml:"servers"`
}

// SetProxy sets the reverse proxy for the backend
func (b *Backend) SetProxy(proxy *httputil.ReverseProxy) {
	b.proxy = proxy
}

// SetAlive sets the alive status for the backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.alive = alive
}

// IsAlive returns whether the backend is alive
func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	defer b.mux.RUnlock()
	return b.alive
}

// GetProxy returns the backend's proxy
func (b *Backend) GetProxy() *httputil.ReverseProxy {
	return b.proxy
}

// GetNextBackend returns the next available backend using round-robin
func (s *Server) GetNextBackend() *Backend {
	s.mux.Lock()
	defer s.mux.Unlock()

	next := s.index
	for i := 0; i < len(s.Backends); i++ {
		idx := (next + i) % len(s.Backends)
		if s.Backends[idx].IsAlive() {
			s.index = (idx + 1) % len(s.Backends)
			return &s.Backends[idx]
		}
	}
	return nil
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
