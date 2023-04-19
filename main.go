package main

import (
	"log"
	"os"

	echo_server "github.com/nitinstp23/protohackers-golang/echo-server"
)

func main() {
	port := "9001"
	maxConnections := 2

	echo_server := echo_server.NewEchoServer(port, maxConnections)
	if err := echo_server.Start(); err != nil {
		log.Printf("failed to start server on port %s, error: %s", port, err)
		os.Exit(1)
	}
}
