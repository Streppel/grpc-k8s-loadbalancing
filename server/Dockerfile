FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/grpc-server ./server/server.go

FROM golang:1.21-alpine
WORKDIR /app
COPY --from=builder /app/grpc-server .
CMD ["./grpc-server"]
