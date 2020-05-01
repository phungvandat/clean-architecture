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

# HOW TO GEN SSL/TLS CERT 
- `cat gen_certs.sh.example > gen_certs.sh`