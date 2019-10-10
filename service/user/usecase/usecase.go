package usecase

import (
	"context"

	"github.com/phungvandat/identity-service/model/request"
	"github.com/phungvandat/identity-service/model/response"
	userService "github.com/phungvandat/identity-service/service/user"
	userRepo "github.com/phungvandat/identity-service/service/user/repository"
	"github.com/phungvandat/identity-service/util/engine"
)

type userUsecase struct {
	userRepo userRepo.Repository
}

func NewUserUseCase(userRepo userRepo.Repository) userService.Service {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (useCase *userUsecase) FindByID(ctx context.Context, req request.FindByID) (*response.FindByID, error) {
	user, err := useCase.userRepo.FindByID(ctx, req.UserID)

	if err != nil {
		return nil, err
	}

	return &response.FindByID{
		User: user,
	}, nil
}

func (useCase *userUsecase) TestAddTranslateQuery(ctx context.Context, req request.TestAddTranslateQuery) (*response.TestAddTranslateQuery, error) {
	query := &engine.Query{}
	res := &response.TestAddTranslateQuery{}

	if req.CreatedAt != nil {
		query.AddFilter("createdAt", engine.GreaterThan, *req.CreatedAt)
	}

	if req.Fullname != "" {
		query.AddFilter("fullname", engine.Equal, req.Fullname)
	}

	if req.Available != nil {
		query.AddFilter("available", engine.Equal, req.Available)
	}

	users, err := useCase.userRepo.TestAddTranslateQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	res.Users = users
	return res, nil
}
