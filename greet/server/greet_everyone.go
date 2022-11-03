package main

import (
	"io"
	"log"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone server was invoked")

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("stream.Recv() got an error - GreetEveryone | error: %+v\n", err)
			return err
		}

		result := "Hello " + msg.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("stream.Send() got an error - GreetEveryone | error: %+v\n", err)
			return err
		}
	}
}
