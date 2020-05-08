package user

import (
	"context"
	"time"

	"github.com/phungvandat/clean-architecture/model/domain"
	repoOptions "github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	"github.com/phungvandat/clean-architecture/repository"
	"github.com/phungvandat/clean-architecture/util/errors"
	"github.com/phungvandat/clean-architecture/util/transaction"
)

type userService struct {
	repo repository.Repository
	txer transaction.Txer
}

// NewUserService func
func NewUserService(repo repository.Repository, txer transaction.Txer) Service {
	return &userService{
		repo: repo,
		txer: txer,
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

// Create function handles logic business
func (s *userService) Create(ctx context.Context, req userReq.Create) (*userRes.Create, error) {
	txPool, err := s.txer.Begin(ctx)
	defer s.txer.Commit(ctx, txPool)

	if err != nil {
		return nil, errors.OccurredError
	}

	user := &domain.User{
		Fullname:  req.Fullname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err = s.repo.User.Create(ctx, user, &repoOptions.RepoOptions{TX: txPool})
	if err != nil {
		s.txer.RollBack(ctx, txPool)
		return nil, err
	}

	return &userRes.Create{
		User: user,
	}, nil
}
