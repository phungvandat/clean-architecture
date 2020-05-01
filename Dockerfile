FROM golang:1.13-alpine as builder

RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates

ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org

WORKDIR /github.com/phungvandat/clean-architecture

COPY . .

RUN go mod tidy
RUN go mod download
RUN go mod verify

RUN go build -o clearn-architecture-go /github.com/phungvandat/clean-architecture/cmd/server/main.go

CMD ["./clearn-architecture-go"]
