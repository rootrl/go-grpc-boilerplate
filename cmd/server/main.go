package main

import (
	"log"
	"net"
	pb "github.com/rootrl/grpc-hello/protos/common"
	"github.com/rootrl/grpc-hello/internal/service" 
	"google.golang.org/grpc"
)


func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	log.Println("Server listeing at", lis.Addr())

	s := grpc.NewServer()

	greetingService := service.NewGreetingServer()

	pb.RegisterGreetingServiceServer(s, greetingService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve: %v", err)
	}

}


