package user

import (
	"context"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	"github.com/stretchr/testify/mock"
)

// UserMock struct
type UserMock struct {
	mock.Mock
}

// FindByID func
func (sm *UserMock) FindByID(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error) {
	args := sm.Called(ctx, req)

	var r0 *userRes.FindByID

	if rf, ok := args.Get(0).(func(context.Context, userReq.FindByID) *userRes.FindByID); ok {
		r0 = rf(ctx, req)
	} else if args.Get(0) != nil {
		r0 = args.Get(0).(*userRes.FindByID)
	}

	var r1 error

	if rf, ok := args.Get(1).(func(context.Context, userReq.FindByID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
