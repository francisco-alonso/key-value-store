# Use official Golang image for building the binary
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the CLI binary
RUN go build -o kvstore ./cmd/kvstorecli

# Use a minimal image for running the binary
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/kvstore .

# Expose a default port (for future API usage)
EXPOSE 8080

# Set the entrypoint
ENTRYPOINT ["./kvstore"]
