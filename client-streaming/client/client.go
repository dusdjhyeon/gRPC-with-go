// client.go
package main

import (
	"context"
	"log"
	"time"

	proto "main/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewClientStreamingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetServerResponse(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	messages := []string{"message #1", "message #2", "message #3", "message #4", "message #5"}

	for _, msg := range messages {
		if err := stream.Send(&proto.Message{Message: msg}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, msg, err)
		}
		log.Printf("[client to server] %s", msg)
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("[server to client] %d", r.GetValue())
}
