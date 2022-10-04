package main

import (
	"context"
	"log"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Hermawan",
	})

	if err != nil {
		log.Fatal("couldn't greet: ", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
