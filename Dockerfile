FROM golang:alpine 

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o  main ./cmd/api/main.go

CMD ["./main"]