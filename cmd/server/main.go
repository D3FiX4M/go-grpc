package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func main() {
	newServer := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.
}
