FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o  main ./cmd/api/main.go

FROM scratch

COPY --from=builder /app/main main

CMD ["./main"]