compile:
	protoc -I=. --go_out=./another-artifact-spec/gen --go-grpc_out=./another-artifact-spec/gen ./another-artifact-spec/proto/api.proto