package user

import (
	"context"

	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	userproto "github.com/phungvandat/clean-architecture/transport/grpc/proto/user"
	"github.com/phungvandat/clean-architecture/util/helper"
)

// FindByID func encode
func FindByID(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*userRes.FindByID)
	user := &userproto.User{}
	err := helper.TransformValue(res.User, user)
	if err != nil {
		return nil, err
	}
	return &userproto.FindByIDResponse{
		User: user,
	}, nil
}
