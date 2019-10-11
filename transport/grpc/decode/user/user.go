package user

import (
	"context"

	userproto "github.com/phungvandat/clean-architecture/transport/grpc/proto/user"
	"github.com/phungvandat/clean-architecture/model/request"
)

func FindByID(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*userproto.FindByIDRequest)
	return request.FindByID{
		UserID: req.UserID,
	}, nil
}
