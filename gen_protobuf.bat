set PROTO_PATH=./internal/infrastructure/delivery/grpc/advertisment
protoc --go_out=%PROTO_PATH% --go-grpc_out=%PROTO_PATH% ./protobuf/advertisment.proto