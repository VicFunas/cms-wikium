[Install golang](https://go.dev/doc/install)

[Install protoc](https://protobuf.dev/installation/#package-manager)

Install golang gRPC plugins:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- Proto directory: Proto service definition (For grpc and generate interfaces)

Generate interface files
```sh
protoc --go_out=. --go-grpc_out=. proto/greeter.proto
```