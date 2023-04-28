# common

## install cmd
* protoc
The easiest way to do this is to download the pre-built binary from the official protobuf release page (https://github.com/protocolbuffers/protobuf/releases). They should download the appropriate binary for their system and extract it to a directory of their choice.
 
* go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
 
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
 
* go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
 
* This command will generate Go code from the auth.proto file and place it in the current directory.
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. auth.proto

