// client.go
package main

import (
	"context"
	"io"
	"log"
	"time"

	proto "main/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewServerStreamingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetServerResponse(ctx, &proto.Number{Value: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetServerResponse(_) = _, %v", c, err)
		}
		log.Printf("[server to client] %s", r.GetMessage())
	}
}
