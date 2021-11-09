package main

import (
	apipb "REST-gRPC-Fibonacci/pkg/api/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, errX := strconv.Atoi(flag.Arg(0))
	if errX != nil {
		log.Fatal(errX)
	}

	y, errY := strconv.Atoi(flag.Arg(1))
	if errY != nil {
		log.Fatal(errY)
	}

	conn, errConn := grpc.Dial(":8080", grpc.WithInsecure())
	if errConn != nil {
		log.Fatal(errConn)
	}

	c := apipb.NewGetFibonacciServiceClient(conn)
	res, errRes := c.GetFibonacci(context.Background(), &apipb.FibonacciRequest{X: int64(x), Y: int64(y)})
	if errRes != nil {
		log.Fatal(errRes)
	}

	log.Println(res.GetResult())

}
