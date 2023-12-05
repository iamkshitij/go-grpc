package main

import (
	pb "go-grpc/calculator/proto"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	//res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
	//	Number: n,
	//})
	//
	//if err != nil {
	//	e, ok := status.FromError(err)
	//}

	//log.Printf("Sqrt :%d\n", res.Result)

}
