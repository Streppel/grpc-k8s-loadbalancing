package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "github.com/Streppel/grpc-k8s-loadbalancing/proto/github.com/Streppel/grpc-k8s-loadbalancing"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("grpc-service:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	for {
		response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "hello from client"})
		if err != nil {
			log.Printf("Error calling SayHello: %v", err)
		} else {
			log.Printf("Response from server in Pod %s: %s", getPodName(), response.Message)
		}

		// Sleep for a random duration between 2 and 5 seconds
		sleepDuration := time.Duration(rand.Intn(4)+2) * time.Second
		time.Sleep(sleepDuration)
	}
}
