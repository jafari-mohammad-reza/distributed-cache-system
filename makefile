.PHONY: protobuf
protobuf: protoc --go_out=. --go-grpc_out=. service.proto