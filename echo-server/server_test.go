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
	inputMessage := "foobar"

	server := echo_server.NewEchoServer(port)
	go server.Start()

	// wait for the server to start
	// TODO: find a better way to do this
	time.Sleep(1000 * time.Millisecond)

	// start TCP client on server port
	conn, _ := net.Dial("tcp", ":"+port)

	fmt.Fprintln(conn, inputMessage)
	serverReply, _ := bufio.NewReader(conn).ReadString('\n')
	serverMessage := strings.ReplaceAll(serverReply, "\n", "")

	if serverMessage != inputMessage {
		t.Errorf("Reply from server %q is not equal to expected %q", serverMessage, inputMessage)
	}
}
