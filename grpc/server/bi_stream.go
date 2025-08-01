package main

import (
	"io"
	"log"

	pb "github.com/Pulkit0707/go/grpc/proto"
)

func (s*helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingClient) error{
	for {
		req,err:=stream.Recv()
		if err==io.EOF{
			return nil
		}
		if err!=nil{
			return err
		}
		log.Printf("Got request with name : %v",req.Name)
		res:=&pb.HelloReponse{
			Message: "Hello"+req.Name,
		}
		if err:=stream.Send(res); err!=nil{
			return err
		}
	}
}