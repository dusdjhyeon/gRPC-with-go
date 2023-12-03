// server.go
package main

import (
	"io"
	"log"
	"net"
	"google.golang.org/grpc"
	proto "main/proto"
)

type server struct {
	proto.BidirectionalServer
}

func (s *server) GetServerResponse(stream proto.Bidirectional_GetServerResponseServer) error {
	log.Printf("Server processing gRPC bidirectional streaming.")
	messages := make([]*proto.Message, 0)
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			for _, msg := range messages {
				if err := stream.Send(msg); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			return err
		}
		messages = append(messages, message)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterBidirectionalServer(s, &server{})
	log.Println("Server is running on port 50051.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
