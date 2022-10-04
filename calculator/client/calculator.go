package main

import (
	"context"
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
