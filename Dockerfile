# Use the official Golang image as the builder
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/app

# Use a minimal image for the final container
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8000

CMD ["/app/main"]
