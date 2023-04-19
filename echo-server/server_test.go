package echo_server_test

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	echo_server "github.com/nitinstp23/protohackers-golang/echo-server"
)

func TestStartWithOneClient(t *testing.T) {
	port := "9001"
	maxConnections := 1
	inputMessage := "foobar"

	server := startServer(port, maxConnections)

	// start TCP client on server port
	conn, _ := net.Dial("tcp", ":"+port)

	fmt.Fprintln(conn, inputMessage)
	serverReply, _ := bufio.NewReader(conn).ReadString('\n')
	serverMessage := strings.ReplaceAll(serverReply, "\n", "")

	if serverMessage != inputMessage {
		t.Errorf("Reply from server %q is not equal to expected %q", serverMessage, inputMessage)
	}

	conn.Close()
	server.Stop()
}

func TestServerMaxConnections(t *testing.T) {
	port := "9001"
	maxConnections := 2

	server := startServer(port, maxConnections)

	// start TCP clients on server port
	conn1, _ := net.Dial("tcp", ":"+port)
	if server.TotalConnections != 1 {
		t.Errorf("totalConnections %d is not equal to expected %d", server.TotalConnections, 1)
	}

	conn2, _ := net.Dial("tcp", ":"+port)
	if server.TotalConnections != 2 {
		t.Errorf("totalConnections %d is not equal to expected %d", server.TotalConnections, 2)
	}

	conn3, _ := net.Dial("tcp", ":"+port)
	if server.TotalConnections != 2 {
		t.Errorf("totalConnections %d is not equal to expected %d", server.TotalConnections, 2)
	}

	conn1.Close()
	conn2.Close()
	conn3.Close()

	server.Stop()
}

func startServer(port string, maxConnections int) *echo_server.EchoServer {
	server := echo_server.NewEchoServer(port, maxConnections)
	go server.Start()

	// wait for the server to start
	// TODO: find a better way to do this
	time.Sleep(1000 * time.Millisecond)

	return server
}
