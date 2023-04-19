package echo_server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

type EchoServer struct {
	port             string
	TotalConnections int
	maxConnections   int
	quit             chan interface{}
	listener         net.Listener
}

var stopMessage = "STOP"

func NewEchoServer(port string, maxConnections int) *EchoServer {
	return &EchoServer{
		port:             port,
		maxConnections:   maxConnections,
		TotalConnections: 0,
		quit:             make(chan interface{}),
	}
}

func (echoServer *EchoServer) Start() error {
	listener, err := net.Listen("tcp", ":"+echoServer.port)
	if err != nil {
		log.Printf("failed to start server on port %s, error: %s", echoServer.port, err)
		return err
	}

	echoServer.listener = listener
	log.Printf("started TCP server on port %s", echoServer.port)

	for {
		select {
		case <-echoServer.quit:
			log.Printf("stopping server on port %s", echoServer.port)
			return nil
		default:
			if echoServer.TotalConnections < echoServer.maxConnections {
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("failed to accept connection on port %s, error: %s", echoServer.port, err)
					return err
				}

				go echoServer.handleNewConnection(conn)
				echoServer.TotalConnections += 1
			}
		}
	}
}

func (echoServer *EchoServer) Stop() {
	log.Printf("stopping server on port %s", echoServer.port)
	echoServer.listener.Close()
	close(echoServer.quit)
}

func (echoServer *EchoServer) handleNewConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("accepting new connection from %s", conn.RemoteAddr())

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			log.Printf("failed to read data, error: %s", err)
			return
		}

		message := strings.TrimSpace(string(netData))

		if message == stopMessage {
			log.Printf("stopping connection: %d", echoServer.TotalConnections)
			echoServer.TotalConnections -= 1
			break
		}

		conn.Write([]byte(netData))
	}
}
