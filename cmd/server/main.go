package main

import (
	apipb "REST-gRPC-Fibonacci/pkg/api/proto"
	"REST-gRPC-Fibonacci/pkg/fibonacci"
	"REST-gRPC-Fibonacci/pkgHTTP"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type RedisCache struct {
	client *redis.Client
}

func main() {

	//Прослушиваем порт
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to listen on port :8080: %v", err)
	}

	//инициализируем сервер gRPC
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &fibonacci.GRPCServer{}

	apipb.RegisterGetFibonacciServiceServer(s, srv)


	//Запускаем сервер gRPC \ Отключаемся командой CTRL+C
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :8080")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	//Подключаем serverHTTP
	pkgHTTP.NewServerHTTP()
	log.Println("Connecting serverHTTP")

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()

	//Закрываем прослушивание порта
	listener.Close()

	log.Println("Done.")
}