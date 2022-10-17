package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/HermawanArifin/grpc-project/calculator/proto"
)

const clientAddress = "localhost:50052"

func main() {
	conn, err := grpc.Dial(clientAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to connect: %+v\n", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	// addition(client)
	// primeStream(client)
	average(client)
}
