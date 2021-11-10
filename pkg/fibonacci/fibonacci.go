package fibonacci

import (
	"REST-gRPC-Fibonacci/pkg/api/proto"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

type GRPCServer struct {
	apipb.UnimplementedGetFibonacciServiceServer
}

// Calculating a number from the Fibonacci sequence
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

// Returns a slice of a sequence of numbers from the Fibonacci series from x to y
func GetFibonacciSlice(x, y int) []int {
	f := fibonacci()
	var result []int
	for i := 0; i <= y; i++ {
		value := f()
		if i >= x && i <= y {
			result = append(result, value)
		}
	}
	return result
}

func (s *GRPCServer) GetFibonacci(ctx context.Context, req *apipb.FibonacciRequest) (*apipb.FibonacciResponse, error) {

	x := int(req.GetX())
	y := int(req.GetY())

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	key := string(x) + " " + string(y)
	fmt.Println(key)

	val1, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")

		slice := GetFibonacciSlice(x, y)
		result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")

		 err := rdb.Set(ctx, key, result, 0).Err()
    	if err != nil {
        	panic(err)
    	}

		return &apipb.FibonacciResponse{Result: result}, nil

    } else if err != nil {
        panic(err)
    } else {
        fmt.Println(key, val1)

		return &apipb.FibonacciResponse{Result: val1}, nil
    }
}