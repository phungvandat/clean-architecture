package usecase

import (
	"context"

	"github.com/phungvandat/identity-service/model/request"
	"github.com/phungvandat/identity-service/model/response"
	userService "github.com/phungvandat/identity-service/service/user"
	userRepo "github.com/phungvandat/identity-service/service/user/repository"
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
