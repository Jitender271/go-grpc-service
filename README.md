# go-grpc-service

# Go gRPC Service with ScyllaDB

This repository contains a microservice written in Go that utilizes gRPC for communication and ScyllaDB as the database. The service manages movie information in a Keyspace named "bookmyshow" with a table named "movies."

## Prerequisites

- Go installed on your machine
- Protocol Buffers (protobuf) compiler installed
- ScyllaDB installed and running

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Jitender271/go-grpc-service.git
   cd go-grpc-service
   
2. Install Dependencies
```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

   
   ```bash
   go mod vendor
      or
   go mod tidy
```

3. Compile Protocol Buffers:
  ```bash
protoc --proto_path= proto/*.proto --go_out=resources/ --go-grpc_out=resources/ --grpc-gateway_out=resources/
```


## ScyllaDB Configuration

1. Connect to your ScyllaDB instance(on Docker) and create the keyspace:
   
  ```bash
CREATE KEYSPACE IF NOT EXISTS bookmyshow WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};
```


2. Create the "movies" table:
   ```bash
   CREATE TABLE IF NOT EXISTS movies (
    movie_id TEXT,
    name TEXT PRIMARY KEY,
    genre TEXT,
    description TEXT,
    rating TEXT,
    );
   ```

## Running the Service

go build

go run main.go

## Usage

Use a gRPC client to interact with the service. See the proto/movie.proto file for service definition.

## Contributing

Feel free to contribute by opening issues or creating pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.



