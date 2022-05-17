package main

import (
	"fmt"
	"log"
	"net"
	"no_vcs/me/example-grpc-service/internal/authors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var defaultPort = 5001

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", defaultPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./ca-cert.pem", "./ca-key.pem")
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	grpcServer := grpc.NewServer(opts...)

	authors.RegisterAuthorsServer(grpcServer, Server{})
	grpcServer.Serve(lis)
}
