package main

import (
	"log"
	"time"

	pb "github.com/Pulkit0707/go/grpc/proto"
)

func (s*helloServer) SayHelloServerStreaming(req*pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error{
	log.Printf("got request with names:%v", req.Names)
	for _,name := range req.Names{
		res:= &pb.HelloReponse{
			Message: "Hello"+name,
		}
		if err:= stream.Send(res); err!=nil{
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}