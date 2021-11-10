package fibonacci

import (
	"REST-gRPC-Fibonacci/pkg/api/proto"
	"context"
	"fmt"
	"log"
	"strings"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/cache/v8"
	"time"
)

type GRPCServer struct {
	apipb.UnimplementedGetFibonacciServiceServer
}

type Object struct {
    slice []int
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

	ring := redis.NewRing(&redis.RingOptions{
        Addrs: map[string]string{
            "server1": ":6379",
            "server2": ":6380",
        },
    })

	cacheSliceFibonacci := cache.New(&cache.Options{
        Redis:      ring,
        LocalCache: cache.NewTinyLFU(1000, time.Minute),
    })
	key := string(x) + "" + string(y)

	var checkSlice []int
	errGetCache := cacheSliceFibonacci.Get(ctx, key, &checkSlice)
	if errGetCache != nil {
		log.Fatal(errGetCache)
	}

	if checkSlice == nil {

		slice := GetFibonacciSlice(x, y)

		if err := cacheSliceFibonacci.Set(&cache.Item{
        	Ctx:   ctx,
       	 	Key:   key,
        	Value: slice,
        	TTL:   time.Hour,
		}); err != nil {
			panic(err)
		}

		result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")

		return &apipb.FibonacciResponse{Result: result}, nil
	}

	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(checkSlice)), ", "), "[]")

	return &apipb.FibonacciResponse{Result: result}, nil
}