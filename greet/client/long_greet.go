package main

import (
	"context"
	"log"
	"time"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	request := []*pb.GreetRequest{
		{FirstName: "Heejin"},
		{FirstName: "Karina"},
		{FirstName: "Sana"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while streaming: %+v\n", err)
	}

	for idx := range request {
		log.Print("Sending request", request[idx])
		stream.Send(request[idx])
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close stream: %+v\n", err)
	}

	log.Printf("Result: %+v\n", res.Result)
}
