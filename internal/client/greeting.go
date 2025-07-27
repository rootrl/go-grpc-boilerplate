package client

import (
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/rootrl/grpc-hello/protos/common"
)

type Client struct{
	Addr string
}

func NewClient() *Client {
	return &Client{
		Addr: "localhost:50051",
	}
}

func (s *Client) Run() {
	serverAddr := s.Addr

	opts := []grpc.DialOption {
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)

	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, opts...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	req := &pb.GetGreetingRequest{
		Id: 123,
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.GetGreeting(ctx, req)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting from server: %s", res.GetGreeting())
}
