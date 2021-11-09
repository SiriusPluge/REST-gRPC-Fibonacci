package main

import (
	apipb "REST-gRPC-Fibonacci/pkg/api/proto"
	"REST-gRPC-Fibonacci/pkg/fibonacci"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	s := grpc.NewServer()
	srv := &fibonacci.GRPCServer{}
	apipb.RegisterGetFibonacciServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if errS := s.Serve(l); errS != nil {
		log.Fatal(errS)
	}

}
