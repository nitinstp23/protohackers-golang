package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	port := "9001"

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("failed to start server on port %s, error: %s", port, err)
		os.Exit(1)
	}

	defer l.Close()
	log.Printf("started TCP server on port %s", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept connection on port %s, error: %s", port, err)
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("accepting new connection from %s", conn.RemoteAddr())

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("failed to read data, error: %s", err)
			return
		}

		conn.Write([]byte(netData))
	}
}
