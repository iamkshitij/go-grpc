package main

import (
	"context"
	pb "go-grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("Do Sum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalf("could not sum %v\n", err)
	}

	log.Printf("Sum :%d\n", res.Result)

}
