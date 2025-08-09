package utils

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Method string
	Path   string
}

func ReadRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	parts := strings.Fields(strings.TrimSpace(line))
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request line")
	}

	return &Request{
		Method: parts[0],
		Path:   parts[1],
	}, nil
}

func BuildHTTPResponse(statusCode int, statusText string, contentType string, body string) string {
	return fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
		statusCode,
		statusText,
		contentType,
		len(body),
		body,
	)
}
