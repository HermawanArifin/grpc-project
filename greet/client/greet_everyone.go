package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("[doGreetEveryone] c.GreetEveryone() got an error - %+v\n", err)
	}

	request := []*pb.GreetRequest{
		{
			FirstName: "Jeon Heejin",
		},
		{
			FirstName: "Nakamura Kazuha",
		},
		{
			FirstName: "Shin Yuna",
		},
	}

	var wg sync.WaitGroup

	// Send the request
	wg.Add(1)
	go func() {
		for idx := range request {
			log.Printf("Sending request: %+v\n", request[idx])
			stream.Send(request[idx])
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
		wg.Done()
	}()

	// Receive the response
	wg.Add(1)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				break
			}

			if err != nil {
				log.Fatalf("[doGreetEveryone] stream.Recv() got an error | err: %+v\n", err)
			}
			fmt.Println("Result:", msg.Result)
		}
	}()
	wg.Wait()
}
