package main

import (
	"context"
	"log"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

func (s *Server) Addition(ctx context.Context, in *pb.CalculatorAdditionRequest) (*pb.CalculatorAdditionResult, error) {
	log.Printf("Addition function was invoked with: %+v\n", in)

	return &pb.CalculatorAdditionResult{
		Result: in.FirstDigit + in.SecondDigit,
	}, nil
}
