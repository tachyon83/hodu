all:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		hodu.proto
	go build -x -o hodu cmd/main.go
