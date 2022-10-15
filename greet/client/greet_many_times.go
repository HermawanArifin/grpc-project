package main

import (
	"context"
	"io"
	"log"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func doStreamGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doStreamGreetManyTimes initiated")

	req := &pb.GreetRequest{
		FirstName: "wibi",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("c.GreetManyTimes() got error: %+v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming: %+v", err)
		}

		log.Printf("message: %s", msg.Result)
	}
}
