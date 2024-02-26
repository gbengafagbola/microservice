FROM golang:1.16

WORKDIR /app

COPY go-service .

RUN go build -o go-service .

CMD ["./go-service"]
