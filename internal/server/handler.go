package server

import (
	"fmt"
	"net"
	"time"

	"github.com/reza-zeinali/golang-client-server/internal/utils"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	request, err := utils.ReadRequest(conn)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	var responseBody string
	switch request.Path {
	case "/ping":
		responseBody = "pong"
	case "/time":
		responseBody = time.Now().Format(time.RFC822)
	default:
		responseBody = "404 not found"
	}

	response := utils.BuildHTTPResponse(200, "OK", "text/plain", responseBody)
	conn.Write([]byte(response))
}
