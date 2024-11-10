package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
    "os"


	"github.com/fatih/color"
    "gopkg.in/yaml.v2"
)

type Backend struct {
    URL string `yaml:"url"`
}

type Server struct {
    Name string `yaml:"name"`
    Port int `yaml:"port"`
    Backends []Backend `yaml:"backends"`
}

type Config struct {
    Servers []Server `yaml:"servers"`
}

func loadConfig(filename string) (Config, error) {
    var config Config

    file, err := os.Open(filename)
    if err != nil {
        return config, err
    }
    defer file.Close()

    decoder := yaml.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        return config, err
    }
    return config, nil
}

func newRover(target string) (*httputil.ReverseProxy, error) {
    url, err := url.Parse(target)
    if err != nil {
        return nil, err
    }
    proxy := httputil.NewSingleHostReverseProxy(url)
    proxy.ModifyResponse = func(resp *http.Response) error {
        return nil
    }
    return proxy, nil
}

func main(){
    config, err := loadConfig("config/config.yml")
	if err != nil {
		color.Set(color.FgRed)
		log.Fatalf("Could not load config: %v", err)
		color.Unset()
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

    for _, server := range config.Servers {
        for _, backend := range server.Backends {
            proxy, err := newRover(backend.URL)
			if err != nil {
    			color.Set(color.FgRed)
	    		log.Fatalf("Could not set up proxy for backend %s: %v", backend.URL, err)
		    	color.Unset()
			}
 			serverName := server.Name
			serverPort := server.Port
			proxyCopy := proxy          
            http.HandleFunc(fmt.Sprintf("/%s", server.Name), func(w http.ResponseWriter, r *http.Request){
                proxyCopy.ServeHTTP(w, r)
            })
            color.Set(color.FgGreen)
            log.Printf("Rover listening for %s on port %d to backend server %s", serverName, serverPort, backend.URL)
		    color.Unset()
        }
    }
    log.Println("Starting rover on :8080")
    color.Unset()
    if err := http.ListenAndServe(":8080", nil); err != nil {
        color.Set(color.FgRed)
        log.Fatalf("Server failed: %v", err)
        color.Unset()
    }

}
