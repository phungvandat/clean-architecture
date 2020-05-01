package user

import (
	"context"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userproto "github.com/phungvandat/clean-architecture/transport/grpc/proto/user"
)

// FindByID func
func FindByID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*userproto.FindByIDRequest)
	return userReq.FindByID{
		UserID: req.UserID,
	}, nil
}
