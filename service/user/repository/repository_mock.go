package repository

import (
	"context"

	"github.com/phungvandat/identity-service/model/domain"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (rm *RepositoryMock) FindByID(ctx context.Context, id string) (*domain.User, error) {
	args := rm.Called(ctx, id)

	var r0 *domain.User
	if rf, ok := args.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
