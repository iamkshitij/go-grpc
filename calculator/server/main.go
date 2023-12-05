package main

import (
	pb "go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while connecting to the server: %v", err)
	}

	log.Printf("Listening to the port %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
