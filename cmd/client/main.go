package main

import (
	"context"
	"fmt"
	"net"
	"no_vcs/me/example-grpc-service/internal/authors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx := context.Background()
	// create a connection with IDP config gRPC service
	idpConfigAddress := net.JoinHostPort("localhost", "5001")

	creds, err := credentials.NewClientTLSFromFile("../../ca-cert.pem", "")
	if err != nil {
		panic(err)
	}

	grpcDialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	idpConn, err := grpc.DialContext(
		ctxTimeout,
		idpConfigAddress,
		grpcDialOptions...,
	)
	if err != nil {
		panic(err)
	}
	defer idpConn.Close()

	ctxTimeout, cancel = context.WithTimeout(ctxTimeout, 5*time.Second)
	defer cancel()

	cli := authors.NewAuthorsClient(idpConn)
	res, err := cli.All(ctxTimeout, &authors.AllFilter{})
	if err != nil {
		panic(err)
	}
	for _, a := range res.Authors {
		fmt.Println(a.Name)
	}
}
