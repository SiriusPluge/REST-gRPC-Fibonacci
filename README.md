# A Fibonacci project
## Services were used:
1. HTTP-server
2. gRPC-server
3. Client
4. Redis

# Description:
Реализовать сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.
Сервис должен отвечать на запросы и возвращать ответ. В ответе должны быть перечислены все числа, последовательности Фибоначчи с порядковыми номерами от x до y.


# To use it, you must enter the following commands:
## Cloning the repository
1. `mkdir $HOME/src/github.com/REST-gRPC-Fibonacci`
2. `cd $HOME/src/github.com/REST-gRPC-Fibonacci`
3. `git clone https://github.com/SiriusPluge/REST-gRPC-Fibonacci.git`
## Docker build:
3. `sudo docker build -t FibonacciServer -f Dockerfile .`
4. `go run cmd/server/main.go`
5. `redis-server`
6. 

## to conduct testing gRPC method, you must:
1.Open two terminals and run:
- get a sequence of Fibonacci numbers: `go run cmd/client/main.go x y`;
2. The result is in the console!

## To conduct testing HTTP method, you must:
1. Open the post man collection, which is located in the root folder of the project and send requests to the HTTP server
2. Открыть коллекцию post man, которая находится в корневой папке проекта и направить запросы на сервер HTTP;

## it is possible to send the following requests:
- get a sequence of Fibonacci numbers/POST: `localhost:8181/fibonacci`;
- add json raw: 
```
- {
    "x": 1,
    "y": 10
}
```

### Congratulations, you have launched the project!!!

## Running the app
- Make sure Redis server is up and running on port `6379`.
- Make sure gRPC server is up and running on port `8080`.
- Make sure HTTP server is up and running on port `8181`.

## P.S.:
created proto:
```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/fibonacci.proto
```