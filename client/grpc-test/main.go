package main

import (
	"context"
	"fmt"
	"log"

	"github.com/phungvandat/clean-architecture/client/grpc-test/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:4001", grpc.WithInsecure())
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
