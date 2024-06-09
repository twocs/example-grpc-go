
# Generate gRPC code
protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		example-grpc-go/example-grpc-go.proto

run-server:
	go run ./server

run-client:
	go run ./client -id=4 -ids="1,3,4,5,77,yomomma"
