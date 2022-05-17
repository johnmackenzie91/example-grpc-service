package main

import (
	"context"
	"fmt"
	"no_vcs/me/example-grpc-service/internal/authors"

	"google.golang.org/grpc/metadata"
)

var _ authors.AuthorsServer = (*Server)(nil)

var svr Server

type Server struct {
	authors.UnimplementedAuthorsServer
}

func init() {
	svr = Server{}
}

func (s Server) All(ctx context.Context, filter *authors.AllFilter) (*authors.AllResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("metadata: %v\n", md)
	}

	return &authors.AllResponse{Authors: []*authors.Author{
		{
			Name: "George Orwell",
		},
	}}, nil
}
