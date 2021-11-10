package pkgHTTP

import (
	"REST-gRPC-Fibonacci/pkg/fibonacci"
	"encoding/json"
	"net/http"
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

	//contentType := req.Header.Get("Content-Type")
	//mediatype, _, err := mime.ParseMediaType(contentType)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//}
	//if mediatype != "application/json" {
	//	http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	//}
	//
	//var jsonData FibVar
	//jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = json.Unmarshal(jsonDataFromHttp, &jsonData)
	//if err != nil {
	//	panic(err)
	//}

	w.Header().Set("Content-Type", "application/json")

	var jsonData FibVar
	err := json.NewDecoder(req.Body).Decode(&jsonData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//result := returningFib(fib, jsonData.X, jsonData.Y)

	slice := fibonacci.GetFibonacciSlice(jsonData.X, jsonData.Y)

	js, err := json.Marshal(slice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
