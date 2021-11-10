FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/server/

CMD ["./main"]