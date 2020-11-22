package env

import (
	"os"
)

// GetHttpPortEnv function get http port from environment
func GetHttpPortEnv() string {
	return os.Getenv("HTTP_PORT")
}

// GetMongoURL function get mongo url from environment
func GetMongoURL() string {
	return os.Getenv("MONGO_URL")
}

// GetMogoDBName function get mongo database name from environment
func GetMogoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}

// GetGrpcPortEnv function get grpc port from environment
func GetGrpcPortEnv() string {
	return os.Getenv("GRPC_PORT")
}

// GetJWTSerectKeyEnv function get jwt serect key from environment
var JWTSerectKeyEnv = getJWTSerectKeyEnv()

func getJWTSerectKeyEnv() string {
	return os.Getenv("JWT_SECRET_KEY")
}

// GetServerKey function get ssl/tls server key from environment
func GetServerKey() string {
	return os.Getenv("SERVER_KEY")
}

// GetServerCRT function get ssl/tls certificate key from environment
func GetServerCRT() string {
	return os.Getenv("SERVER_CRT")
}
