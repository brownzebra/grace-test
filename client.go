package main

import (
	"context"
	"log"
	"time"

	pb "github.com/brownzebra/grpc-test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallSendMsgSize(1024*1024*20), // Allow up to 20 MB
	))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTestServiceClient(conn)

	for size := 1 * 1024 * 1024; size <= 25*1024*1024; size += 1 * 1024 * 1024 {
		payload := make([]byte, size)
		log.Printf("Sending payload of size: %d bytes", size)

		_, err := client.SendMessage(context.Background(), &pb.TestRequest{Payload: payload})
		if err != nil {
			log.Printf("Failed at size %d bytes: %v", size, err)
			break
		} else {
			log.Printf("Success at size %d bytes", size)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
