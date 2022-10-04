package main

import (
	"log"
	"net"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
	"google.golang.org/grpc"
)

const grpcCalculator = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", grpcCalculator)
	if err != nil {
		log.Fatalf("failed to listen on: %+v\n", err)
	}

	log.Printf("listening on: %+v\n", grpcCalculator)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}
}
