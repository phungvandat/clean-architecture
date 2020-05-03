package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	FindByIDFunc func(ctx context.Context, id string) (*domain.User, error)
	FindFunc     func(ctx context.Context, conditions userInput.FindConditions) ([]*domain.User, error)
}

// FindByID function
func (rm *RepositoryMock) FindByID(ctx context.Context, id string) (*domain.User, error) {
	if rm.FindByIDFunc == nil {
		panic("RepositoryMock not declare FindByID function")
	}
	return rm.FindByIDFunc(ctx, id)
}

// Find function
func (rm *RepositoryMock) Find(ctx context.Context, conditions userInput.FindConditions) ([]*domain.User, error) {
	if rm.FindFunc == nil {
		panic("RepositoryMock not declare Find function")
	}
	return rm.FindFunc(ctx, conditions)
}
