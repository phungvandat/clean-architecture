dev: 
	ENV=local go run cmd/server/main.go

test:
	go test ./...
