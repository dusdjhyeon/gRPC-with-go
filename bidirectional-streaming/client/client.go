// client.go
package main

import (
	"context"
	"io"
	"log"
	"time"
	"google.golang.org/grpc"
	proto "main/proto"// proto 파일의 실제 위치로 변경해야 합니다.
)

func makeMessage(message string) *proto.Message {
	return &proto.Message{Message: message}
}

func generateMessages() []*proto.Message {
	messages := []*proto.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
	return messages
}

func sendMessage(client proto.BidirectionalClient) {
	stream, err := client.GetServerResponse(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
			}
			log.Printf("[server to client] %s", in.Message)
		}
	}()
	for _, msg := range generateMessages() {
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
		log.Printf("[client to server] %s", msg.Message)
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewBidirectionalClient(conn)
	sendMessage(client)
}
