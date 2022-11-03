package main

import (
	"io"
	"log"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

func (s *Server) Maximum(stream pb.CalculatorService_MaximumServer) error {
	log.Println("Maximum invoked")

	maxNumber := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("[Maximum] stream.Recv() got error: %+v\n", err)
		}

		if msg.Number < int64(maxNumber) {
			continue
		}

		maxNumber = int(msg.Number)
		err = stream.Send(&pb.CalculatorMaximumResult{
			Result: int64(maxNumber),
		})
		if err != nil {
			log.Fatalf("[Maximum] stream.Send() got error: %+v\n", err)
			return err
		}
	}

	return nil
}
