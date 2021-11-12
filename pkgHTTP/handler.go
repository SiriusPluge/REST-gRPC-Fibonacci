package pkgHTTP

import (
	"REST-gRPC-Fibonacci/pkg/fibonacci"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
)

type FibVar struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//// непосредственно функция подсчета фибоначи.
//func fib(n int) uint {
//	if n < 2 {
//		return 1
//	}
//	return fib(n-2) + fib(n-1)
//}
//
//type fibfunc func(int) uint
//
//// заполнение слайса результирующими значениями
//func returningFib(fib fibfunc, a, b int) []uint {
//	slc := make([]uint, 0, b-a+1)
//	for i := a; i < b; i++ {
//		slc = append(slc, fib(i))
//	}
//	return slc
//}

func GetFibonacci(w http.ResponseWriter, req *http.Request)  {

	w.Header().Set("Content-Type", "application/json")

	var jsonData FibVar
	err := json.NewDecoder(req.Body).Decode(&jsonData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	key := string(jsonData.X) + " " + string(jsonData.Y)
	ctx := context.Background()

	val1, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")

		slice := fibonacci.GetFibonacciSlice(int64(jsonData.X), int64(jsonData.Y))
		result := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")

		 err := rdb.Set(ctx, key, result, 0).Err()
    	if err != nil {
        	panic(err)
    	}

		js, err := json.Marshal(slice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

    } else if err != nil {
        panic(err)
    } else {
        fmt.Println(key, val1)

		js, err := json.Marshal(val1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
    }
}
