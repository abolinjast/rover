# Rover - Simple Reverse Proxy in Go

Rover is a lightweight reverse proxy tool built in Golang. It forwards incoming requests to specified backend servers and relays the responses back to the client. This simple implementation allows for easy customization and is great for learning or as a foundation for more advanced proxy projects.

## Features

- Reverse proxy for forwarding HTTP requests
- Modular design for easy customization
- Color-coded logs for enhanced readability in the terminal
- ASCII banner for a personalized startup display

## Getting Started

### Prerequisites

- Go (version 1.16 or later is recommended)
- Terminal that supports color output (optional, for color-coded logs)

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/abolinjast/rover.git
    cd rover
    ```

2. **Build the project**:

    ```bash
    go build -o rover
    ```

3. **Run Rover**:

    ```bash
    ./rover
    ```

By default, Rover will forward requests based on the configuration provided in the `config/config.yml` file. You can modify the target URLs and servers by editing this configuration file.

### Usage

To start the reverse proxy without building an executable, simply run:

```bash
go run main.go

Once running, Rover will listen on http://localhost:8080 and forward incoming requests to the backend server(s) specified in the configuration file.
Configuration

The target backend URLs are configured in the config/config.yml file. Edit this file to set up the backend servers you want Rover to forward requests to.

Example config/config.yml configuration:

servers:
  - name: server1
    port: 8080
    backends:
      - url: "http://backend1.com"
      - url: "http://backend2.com"

In this example, Rover will route incoming requests for server1 to either http://backend1.com or http://backend2.com based on your configuration. You can add multiple servers and backends as needed.
Project Structure

rover/
├── main.go            # Main application file
├── config/            # Configuration files (e.g., config/config.yml)
├── handlers/          # Custom handlers (future feature)
├── README.md          # Documentation
└── go.mod             # Dependency management

This structure allows you to organize your code and add custom handlers or features as the project grows.

License      
This project is licensed under the MIT License - see the LICENSE file for details.

Contributing 
Contributions are welcome! Feel free to submit a pull request or open an issue if you have ideas for improvement.
