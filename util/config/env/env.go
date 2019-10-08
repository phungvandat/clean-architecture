package env

import (
	"os"
)

func GetPortEnv() string {
	return os.Getenv("PORT")
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetMogoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}

func GetGRPCPortEnv() string {
	return os.Getenv("GRPC_PORT")
}
