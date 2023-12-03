// server.go
package main

import (
	"io"
	"log"
	"net"

	proto "main/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedClientStreamingServer
}

func (s *server) GetServerResponse(stream proto.ClientStreaming_GetServerResponseServer) error {
	log.Printf("Server processing gRPC client-streaming.")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.Number{Value: int32(count)})
		}
		if err != nil {
			return err
		}
		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterClientStreamingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
