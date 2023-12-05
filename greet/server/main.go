package main

import (
	pb "go-grpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error in connection: ", err)
	}
	log.Printf("Listening on port %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Fail to serve: %v\n", err)
	}
}
