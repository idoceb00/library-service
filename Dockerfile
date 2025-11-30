FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o library-service ./cmd/library-service

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/library-service .

EXPOSE 8080
CMD ["./library-service"]