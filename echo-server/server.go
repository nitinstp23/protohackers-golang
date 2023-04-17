package echo_server

import (
	"bufio"
	"log"
	"net"
)

type TCPServer struct {
	port string
}

func NewTCPServer(port string) *TCPServer {
	return &TCPServer{port: port}
}

func (tcp *TCPServer) Start() error {
	l, err := net.Listen("tcp", ":"+tcp.port)
	if err != nil {
		log.Printf("failed to start server on port %s, error: %s", tcp.port, err)
		return err
	}

	defer l.Close()
	log.Printf("started TCP server on port %s", tcp.port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept connection on port %s, error: %s", tcp.port, err)
			return err
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
