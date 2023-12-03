// server.go
package main

import (
	"log"
	"net"

	proto "main/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.ServerStreamingServer
}

func (s *server) GetServerResponse(req *proto.Number, stream proto.ServerStreaming_GetServerResponseServer) error {
	log.Printf("Server processing gRPC server-streaming.")
	messages := []string{"message #1", "message #2", "message #3", "message #4", "message #5"}

	for _, msg := range messages {
		if err := stream.Send(&proto.Message{Message: msg}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterServerStreamingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
