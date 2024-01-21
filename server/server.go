package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Streppel/grpc-k8s-loadbalancing/proto/github.com/Streppel/grpc-k8s-loadbalancing"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer // Embed the UnimplementedGreeterServer struct
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	podName := getPodName()
	fmt.Printf("Received request in Pod %s: %v\n", podName, request.Name)
	message := "Serving from " + podName + " IP " + os.Getenv("MY_POD_IP")
	response := &pb.HelloReply{Message: message}
	return response, nil
}

func getPodName() string {
	podName, found := os.LookupEnv("MY_POD_NAME")
	if !found {
		return "unknown"
	}
	return podName
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
