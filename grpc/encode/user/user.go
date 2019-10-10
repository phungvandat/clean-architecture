package user

import (
	"context"

	userproto "github.com/phungvandat/clean-architecture/grpc/proto/user"
	"github.com/phungvandat/clean-architecture/model/response"
)

func FindByID(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*response.FindByID)
	return &userproto.FindByIDResponse{
		User: &userproto.User{
			Id:       res.User.ID.Hex(),
			Username: res.User.Username,
			Fullname: res.User.Fullname,
			Email:    res.User.Email,
			Phone:    res.User.Phone,
			Role:     res.User.Role,
		},
	}, nil
}
