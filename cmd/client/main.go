package main

import(
	"github.com/rootrl/grpc-hello/internal/client"
)

func main() {
	client := client.NewClient()

	client.Run()
}
