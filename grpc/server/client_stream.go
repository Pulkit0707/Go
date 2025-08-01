package main

import (
	"io"
	"log"

	pb "github.com/Pulkit0707/go/grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientSteeamingServer) error {
	var messages []string
	for{
		req,err:=stream.Recv()
		if err==io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err!=nil{
			return err
		}
		log.Printf("Got request with names:%v",req.Name)
		messages=append(messages, "Hello",req.Name)
	}
}