# REST-gRPC-Fibonacci service

created proto:
```
    protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/fibonachi.proto
```
