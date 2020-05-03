package user

import (
	"context"

	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	"github.com/phungvandat/clean-architecture/repository"
)

type userService struct {
	repo repository.Repository
}

// NewUserService func
func NewUserService(repo repository.Repository) Service {
	return &userService{
		repo: repo,
	}
}

// FindByID function handles logic business
func (s *userService) FindByID(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error) {
	user, err := s.repo.User.FindByID(ctx, req.UserID)

	if err != nil {
		return nil, err
	}

	return &userRes.FindByID{
		User: user,
	}, nil
}

// Find function handles logic business
func (s *userService) Find(ctx context.Context, req userReq.Find) (*userRes.Find, error) {
	conditions := userInput.FindConditions{}

	if req.Fullname != "" {
		conditions.Fullname = req.Fullname
	}

	users, err := s.repo.User.Find(ctx, conditions)
	if err != nil {
		return nil, err
	}

	res := &userRes.Find{
		Users: users,
	}
	return res, nil
}
