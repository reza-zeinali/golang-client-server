package server

import (
	"fmt"
	"net"
)

type Server struct {
	Port string
}

func New(port string) *Server {
	return &Server{Port: port}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
