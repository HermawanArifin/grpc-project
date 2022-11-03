package main

import (
	"log"
	"net"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, key)
		if err != nil {
			log.Fatalf("failed to load credentials: %+v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v\n", err)
	}
}
