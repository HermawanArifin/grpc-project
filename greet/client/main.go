package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/HermawanArifin/grpc-project/greet/proto"
)

const clientAddress = "localhost:50051"

func main() {
	opts := []grpc.DialOption{}
	tls := true
	if tls {
		certFile := "ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("failed to load credentials: %+v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(clientAddress, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to connect: %+v\n", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	doStreamGreetManyTimes(client)
	doLongGreet(client)
	doGreetEveryone(client)
}
