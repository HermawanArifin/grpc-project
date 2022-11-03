package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

func addition(c pb.CalculatorServiceClient) {
	log.Println("addition invoked")

	res, err := c.Addition(context.Background(), &pb.CalculatorAdditionRequest{
		FirstDigit:  10,
		SecondDigit: 20,
	})

	if err != nil {
		log.Fatal("couldn't add: ", err)
	}

	log.Printf("Result : %d", res.Result)
}

func primeStream(c pb.CalculatorServiceClient) {
	log.Println("primeStream invoked")

	res, err := c.PrimeStream(context.Background(), &pb.CalculatorPrimeRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatal("couldn't stream", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming: %+v", err)
		}

		log.Printf("result: %d", msg.Number)
	}
}

func average(c pb.CalculatorServiceClient) {
	inputs := []*pb.CalculatorAverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal("error while streaming average", err)
	}

	for _, input := range inputs {
		stream.Send(input)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error while closing stream", err)
	}

	log.Print("result", res)
}

func maximum(c pb.CalculatorServiceClient) {
	stream, err := c.Maximum(context.Background())
	if err != nil {
		log.Fatalf("[maximum] c.Maximum() got error: %+v\n", err)
	}

	inputs := []*pb.CalculatorMaximumRequest{
		{Number: 100},
		{Number: 4},
		{Number: 101},
		{Number: 12},
		{Number: 15},
		{Number: 1},
	}

	var wg sync.WaitGroup
	// send message
	wg.Add(1)
	go func() {
		for _, input := range inputs {
			err := stream.Send(&pb.CalculatorMaximumRequest{
				Number: input.Number,
			})
			if err != nil {
				log.Printf("[maximum] stream.Send() got error: %+v\n", err)
			}

			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
		wg.Done()
	}()

	// receive message
	wg.Add(1)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				break
			}
			if err != nil {
				log.Printf("[maximum] stream.Recv() got error: %+v\n", err)
			}

			resultStr := strconv.FormatInt(msg.Result, 10)
			log.Println("Maximum number is: ", resultStr)
		}
	}()
	wg.Wait()
}
