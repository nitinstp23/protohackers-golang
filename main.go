package main

import (
	"log"
	"os"

	echo_server "github.com/nitinstp23/protohackers-golang/echo-server"
)

func main() {
	port := "9001"

	tcp_server := echo_server.NewTCPServer(port)
	if err := tcp_server.Start(); err != nil {
		log.Printf("failed to start server on port %s, error: %s", port, err)
		os.Exit(1)
	}
}
