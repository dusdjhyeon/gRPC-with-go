// client.go
package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
    proto "main/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()

    c := proto.NewMyServiceClient(conn)

    response, err := c.MyFunction(context.Background(), &proto.MyNumber{Value: 4})
    if err != nil {
        log.Fatalf("Could not call MyFunction: %v", err)
    }
    log.Printf("gRPC result: %v", response.GetValue())
}
