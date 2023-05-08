
all: go_out

go_out: 
	protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. pkg/proto/*/*.proto
	cp -r github.com/dakong-yi/common/pkg/pb/ pkg
	rm -r github.com