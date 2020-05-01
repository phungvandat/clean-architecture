init:
	cat .env.example > .env
	go mod init
	
dev:
	go mod tidy
	ENV=local go run cmd/server/main.go

test:
	go test ./...

tls-secure:
	go mod tidy
	ENV=tls-secure go run cmd/server/main.go

docker-image:
	docker build -t clearn-architecture-go-image .