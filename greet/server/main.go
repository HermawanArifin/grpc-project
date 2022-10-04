package main

import (
	"log"
	"net"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
	"google.golang.org/grpc"
)

const grpcAddress = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen on: %+v\n", err)
	}

	log.Printf("listening on: %+v\n", grpcAddress)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}
}
