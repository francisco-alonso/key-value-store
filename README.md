# Prerequisites
Before running the project locally or inside a Docker container, make sure you have the following installed:

- Go 1.18+
- Docker (for building and running the container)

# Install Dependencies
The project uses Go modules to manage dependencies. If you haven't already, initialize the Go modules in the project directory (if it's not already done):

```
go mod tidy
```

# Build the Go Project
To build the project locally, run:
```
go build -o kvstore ./cmd/kvstorecli
```
This will compile the Go application and create a binary named kvstore.

# Run the Project Locally
To run the application locally, start the HTTP server using:

```
go run ./cmd/kvstorecli
```
This will launch the HTTP server on localhost:8080. You can now interact with the API using the following endpoints:

```
POST /set: Set a key-value pair (e.g., {"key": "foo", "value": "bar"})
GET /get: Retrieve the value for a key (e.g., /get?key=foo)
DELETE /delete: Delete a key-value pair (e.g., /delete?key=foo)
GET /exists: Check if a key exists (e.g., /exists?key=foo)

```
