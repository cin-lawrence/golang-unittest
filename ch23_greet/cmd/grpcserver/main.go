package main

import (
	"log"
	"net"

	"ch23/greet/adapters/grpcserver"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &grpcserver.GreetServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
