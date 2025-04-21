package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"

	"github.com/abolinjast/rover/internal/config"
	"github.com/abolinjast/rover/internal/proxy"
)

func main() {
	config, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		color.Set(color.FgRed)
		log.Fatalf("Could not load config: %v", err)
	}

	roverAscii := `
 ____   _____     _______ ____
|  _ \ / _ \ \   / / ____|  _ \
| |_) | | | \ \ / /|  _| | |_) |
|  _ <| |_| |\ V / | |___|  _ <
|_| \_\\___/  \_/  |_____|_| \_\

    `
	color.Set(color.FgGreen)
	fmt.Printf("%s\n", roverAscii)
	fmt.Println("Rover is ready to handle requests...")

	// Initialize servers and backends
	for i := range config.Servers {
		server := &config.Servers[i]
		if err := proxy.InitBackends(server); err != nil {
			color.Set(color.FgRed)
			log.Fatalf("Could not initialize backends for %s: %v", server.Name, err)
		}

		// Register handler once per server
		http.HandleFunc(fmt.Sprintf("/%s/", server.Name), func(w http.ResponseWriter, r *http.Request) {
			backend := server.GetNextBackend()
			if backend == nil {
				http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
				return
			}
			backend.GetProxy().ServeHTTP(w, r)
		})

		color.Set(color.FgGreen)
		log.Printf("Rover listening for %s on port %d", server.Name, server.Port)
	}

	log.Println("Starting rover on :8080")
	color.Unset()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		color.Set(color.FgRed)
		log.Fatalf("Server failed: %v", err)
	}
}
