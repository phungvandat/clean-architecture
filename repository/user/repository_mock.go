package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	FindByIDFunc func(ctx context.Context, id string) (*domain.User, error)
}

// FindByID func
func (rm *RepositoryMock) FindByID(ctx context.Context, id string) (*domain.User, error) {
	if rm.FindByIDFunc == nil {
		panic("RepositoryMock not declare FindByID function")
	}
	return rm.FindByIDFunc(ctx, id)
}
