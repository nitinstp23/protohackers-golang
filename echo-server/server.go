package echo_server

import (
	"bufio"
	"log"
	"net"
)

type EchoServer struct {
	port string
}

func NewEchoServer(port string) *EchoServer {
	return &EchoServer{port: port}
}

func (echoServer *EchoServer) Start() error {
	l, err := net.Listen("tcp", ":"+echoServer.port)
	if err != nil {
		log.Printf("failed to start server on port %s, error: %s", echoServer.port, err)
		return err
	}

	defer l.Close()
	log.Printf("started TCP server on port %s", echoServer.port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept connection on port %s, error: %s", echoServer.port, err)
			return err
		}

		go echoServer.handleNewConnection(conn)
	}
}

func (echoServer *EchoServer) handleNewConnection(conn net.Conn) {
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
