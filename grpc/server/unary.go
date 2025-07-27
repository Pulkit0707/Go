package main

import (
	"context"
	pb "github.com/Pulkit0707/go/grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context,req*pb.NoParam)(*pb.HelloReponse,error){
	return &pb.HelloReponse{
		Message: "Hello from server",
	}, nil
}