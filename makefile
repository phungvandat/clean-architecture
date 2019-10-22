init:
	cat .env.example > .env
	go mod init
	
dev:
	go mod tidy
	ENV=local go run cmd/server/main.go

test:
	go test ./...

secure-grpc:
	go mod tidy
	ENV=secure-grpc go run cmd/server/main.go