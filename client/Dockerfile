FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/grpc-client ./client/client.go

FROM golang:1.21-alpine
WORKDIR /app
COPY --from=builder /app/grpc-client .
CMD ["./grpc-client"]
