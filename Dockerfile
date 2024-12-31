# Start from an official Go image for building
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Build the CLI application
RUN go build -o kvstore ./cmd/kvstorecli

# Use a minimal image for the runtime
FROM debian:bullseye-slim

# Copy the compiled binary from the builder stage
COPY --from=builder /app/kvstore /usr/local/bin/kvstore

# Set the entrypoint
ENTRYPOINT ["kvstore"]
