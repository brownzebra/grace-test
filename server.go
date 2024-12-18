package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brownzebra/grpc-test/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Received message of size: %d bytes", len(req.Payload))
	return &pb.TestResponse{Message: "Received"}, nil
}

func main() {
	// gRPC server options to handle large messages
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 20), // Allow up to 20 MB
	}

	s := grpc.NewServer(opts...)
	pb.RegisterTestServiceServer(s, &server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
