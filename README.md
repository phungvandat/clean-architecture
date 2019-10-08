# identity-service

# HOW TO RUN
`cat .env.example > .env`
`make dev`

# HOW TO TEST
`make test`

# HOW TO GEN PROTOC
`protoc --go_out=plugins=grpc:. grpc/proto/user/user.proto`