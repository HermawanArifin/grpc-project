package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	var temp int
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(temp, count)
			return stream.SendAndClose(&pb.CalculatorAverageResult{
				Result: float32(temp) / float32(count),
			})
		}

		if err != nil {
			log.Fatalf("error while streaming Average: %+v\n", err)
		}

		temp += int(req.Number)
		count++
	}
}
