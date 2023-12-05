![client-streaming](https://github.com/dusdjhyeon/gRPC-with-go/assets/73868703/d0e30bfa-6961-4755-94c1-f590689043bc)# gRPC-with-go

There is a source code that implements the four communication methods of gRPC in Golang.
- Unary gRPC (no streaming)
- Bi-directional streaming
- Client to Server streaming
- Server to Client streaming

To run, you can run the following commands in the server and client folders under each folder.
```bash
$ go run server.go
$ go run client.go
```

## Execution result
### Unary gRPC
![unary](https://github.com/dusdjhyeon/gRPC-with-go/assets/73868703/fa75c99c-9ecd-4d07-b977-aefb684195ac)


### Bi-directional streaming
![bidirectional](https://github.com/dusdjhyeon/gRPC-with-go/assets/73868703/ae9a712c-285e-475f-aa74-096b51344d76)


### Client streaming
![client-streaming](https://github.com/dusdjhyeon/gRPC-with-go/assets/73868703/2a75014a-41a6-4172-a04e-75cd89d1caf1)


### Server streaming
![server-streaming](https://github.com/dusdjhyeon/gRPC-with-go/assets/73868703/496a0a8e-ddad-40e4-9863-1f1ff9162f39)


## References
- https://grpc.io/docs/languages/go/basics/
- https://github.com/grpc/grpc-go
- https://protobuf.dev/
