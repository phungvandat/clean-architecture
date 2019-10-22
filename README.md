# CLEAN ARCHITECTURE

![clean architecture](https://raw.githubusercontent.com/phungvandat/clean-architecture/dev/images/clean-arch.png)

# HOW TO SETUP DB
- `docker-compose -f docker-compose-local.yaml up -d`

# HOW TO RUN
### First
- `make init`
### Next
- `make dev`

# HOW TO TEST
- `make test`

# HOW TO GEN PROTOC
- `protoc --go_out=plugins=grpc:. grpc/proto/user/user.proto`

# HOW TO GEN CERT 
- `openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout cert/server.key -out cert/server.pem`