package main

import (
	"log"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

func (s *Server) PrimeStream(in *pb.CalculatorPrimeRequest, stream pb.CalculatorService_PrimeStreamServer) error {
	log.Println("Prime stream was invoked")

	number := in.Number
	k := 2

	for number > 1 {
		if number%2 == 0 {
			stream.Send(&pb.CalculatorPrimeResult{
				Number: int64(k),
			})
			number /= int64(k)
		} else {
			k++
		}
	}

	return nil
}
