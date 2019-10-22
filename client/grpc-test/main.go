package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phungvandat/clean-architecture/client/grpc-test/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

	var dialOption grpc.DialOption = grpc.WithInsecure()

	// Create the client TLS credentials
	if os.Getenv("ENV") == "secure-grpc" {
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM([]byte(os.Getenv("CA_PEM"))) {
			fmt.Errorf("credentials: failed to append certificates")
		}
		creds := credentials.NewTLS(&tls.Config{ServerName: os.Getenv("GRPC_SERVER_ADDR"), RootCAs: cp})
		dialOption = grpc.WithTransportCredentials(creds)
	}

	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER_ADDR"), dialOption)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	identitySrv := user.NewUserServiceClient(conn)
	req := &user.FindByIDRequest{
		UserID: "5c1236498eda730c3516b07d",
	}

	ctx := context.TODO()

	mdCtx := metadata.NewOutgoingContext(ctx, newMetadata(
		map[string]string{
			"myKey": "go.io",
		},
	))
	res, err := identitySrv.FindByID(mdCtx, req)

	if err != nil {
		log.Fatalf("Fail get user by xxx: %v", err.Error())
		return
	}

	fmt.Println("Name: ", res.User.Username)

}

func newMetadata(data map[string]string) metadata.MD {
	return metadata.New(data)
}
