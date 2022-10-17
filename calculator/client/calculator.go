package main

import (
	"context"
	"io"
	"log"

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
