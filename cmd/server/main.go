package main

import (
	"fmt"
	"log"
	"os"

	"golang-client-server/internal/server"
)

func main() {
	port := ":8080"
	s := server.New(port)

	fmt.Printf("Server starting on %s\n", port)
	if err := s.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
		os.Exit(1)
	}
}
