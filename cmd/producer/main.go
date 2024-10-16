package main

import (
	"context"
	"log"
	"time"

	pb "queueapi/pkg/api" // Import the generated package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewQueueClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		sendResp, err := c.SendMessage(ctx, &pb.SendMessageRequest{MessageBody: "Hello, SQS!"})
		if err != nil {
			log.Fatalf("Could not send message: %v", err)
			break
		}
		log.Printf("Message sent, ID: %s", sendResp.MessageId)

		time.Sleep(2 * time.Second)
	}
}
