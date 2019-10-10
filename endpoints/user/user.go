package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	userReq "github.com/phungvandat/identity-service/model/request"
	"github.com/phungvandat/identity-service/service"
)

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

func MakeTestAddTranslateQuery(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userReq.TestAddTranslateQuery)
		res, err := s.UserService.TestAddTranslateQuery(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
