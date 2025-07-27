package service

import (
	"context"
	"fmt"
	pb "github.com/rootrl/grpc-hello/protos/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreetingServer struct {
	pb.UnimplementedGreetingServiceServer
}

func NewGreetingServer() *GreetingServer {
	return &GreetingServer {}
}

func (s *GreetingServer) GetGreeting(ctx context.Context, req *pb.GetGreetingRequest) (*pb.GetGreetingResponse, error) {
	if req.GetId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID must be greater than 0");
	}

	fmt.Println("New message from client: %d", req.GetId())

	greetingMessage := fmt.Sprintf("Hello, client with ID %d!", req.GetId())

	response := &pb.GetGreetingResponse {
		Greeting: greetingMessage,
	}

	return response, nil
}

