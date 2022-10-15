package main

import (
	"fmt"
	"log"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("greet many times was invoked: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i+1)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
