package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	"github.com/phungvandat/clean-architecture/service"
)

// MakeFindByIDEndpoint function return endpoint of find user by id
func MakeFindByIDEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq.FindByID)
		res, err := s.UserService.FindByID(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

// MakeFindEndpoint function return endpoint of find user
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq.Find)
		res, err := s.UserService.Find(ctx, req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
