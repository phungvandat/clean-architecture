package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	FindByIDFunc func(ctx context.Context, id string, options ...*repository.RepoOptions) (*domain.User, error)
	FindFunc     func(ctx context.Context, conditions userInput.FindConditions, options ...*repository.RepoOptions) ([]*domain.User, error)
	CreateFunc   func(ctx context.Context, user *domain.User, options ...*repository.RepoOptions) (*domain.User, error)
}

// FindByID function
func (rm *RepositoryMock) FindByID(ctx context.Context, id string, options ...*repository.RepoOptions) (*domain.User, error) {
	if rm.FindByIDFunc == nil {
		panic("RepositoryMock not declare FindByID function")
	}
	return rm.FindByIDFunc(ctx, id, options...)
}

// Find function
func (rm *RepositoryMock) Find(ctx context.Context, conditions userInput.FindConditions, options ...*repository.RepoOptions) ([]*domain.User, error) {
	if rm.FindFunc == nil {
		panic("RepositoryMock not declare Find function")
	}
	return rm.FindFunc(ctx, conditions, options...)
}

// Create function
func (rm *RepositoryMock) Create(ctx context.Context, user *domain.User, options ...*repository.RepoOptions) (*domain.User, error) {
	if rm.CreateFunc == nil {
		panic("RepositoryMock not declare Create function")
	}
	return rm.CreateFunc(ctx, user, options...)
}
