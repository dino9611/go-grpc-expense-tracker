xprotos:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/*.proto

run-auth:
	cd services/auth && nodemon --exec go run ./cmd/grpc/main.go --signal SIGTERM

run-gateway:
	cd gateway && nodemon --exec go run ./cmd/http/main.go --signal SIGTERM