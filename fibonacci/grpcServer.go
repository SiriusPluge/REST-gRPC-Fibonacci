package fibonacci

import (
	apipb "REST-gRPC-Fibonacci/api/proto"
	"context"
	"fmt"
	"strings"
)

type GRPCServer struct {}

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
func getFibonacciSlice(x, y int) []int {
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
	slice := getFibonacciSlice(x, y)
	result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")
	//fmt.Println(x, y, slice, result)

	return &apipb.FibonacciResponse{Result: result}, nil
}