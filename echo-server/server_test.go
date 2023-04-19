package echo_server_test

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"

	echo_server "github.com/nitinstp23/protohackers-golang/echo-server"
)

func TestStartWithOneClient(t *testing.T) {
	port := "9001"
	inputMessage := "foobar"

	conn, _ := net.Dial("tcp", ":"+port)
	startTCPServer(port)

	fmt.Fprint(conn, inputMessage)
	message, _ := bufio.NewReader(conn).ReadString('\n')

	if message != inputMessage {
		t.Errorf("Reply from server %q is not equal to expected %q", message, inputMessage)
	}
}

func startTCPServer(port string) *echo_server.EchoServer {
	server := echo_server.NewEchoServer(port)
	if err := server.Start(); err != nil {
		log.Printf("failed to start server on port %s, error: %s", port, err)
		return nil
	}

	return server
}
