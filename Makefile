gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:./ --go-grpc_out=:./
	#protoc -I=proto --go_out=./ proto/*.proto
	go mod tidy

clean:
	rm -rf  pb/*
	go clean -cache

run:
	go mod tidy
	go run main.go

test:
	go test -cover -race ./...

server:
	go run server/main.go -port 8080

client:
	go run ./client/main.go -address 0.0.0.0:8080