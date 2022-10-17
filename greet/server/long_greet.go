package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	var res string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("stream.Recv() got error - LongGreet", err)
		}
		res += fmt.Sprintf("Hello, %s!\n", req.FirstName)
	}
}
