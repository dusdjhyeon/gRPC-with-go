package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
	pb "main/grpc"
	proto "main/proto"
)

type server struct {
    proto.MyServiceServer
}

func (s *server) MyFunction(ctx context.Context, in *proto.MyNumber) (*proto.MyNumber, error) {
	result := pb.MyFunc(in.Value)
    return &proto.MyNumber{Value: result}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    proto.RegisterMyServiceServer(s, &server{})
    log.Println("Server is running on port 50051.")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
