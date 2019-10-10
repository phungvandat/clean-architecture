package usecase

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/request"
	"github.com/phungvandat/clean-architecture/model/response"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (um *UsecaseMock) FindByID(ctx context.Context, req request.FindByID) (*response.FindByID, error) {
	args := um.Called(ctx, req)

	var r0 *response.FindByID

	if rf, ok := args.Get(0).(func(context.Context, request.FindByID) *response.FindByID); ok {
		r0 = rf(ctx, req)
	} else if args.Get(0) != nil {
		r0 = args.Get(0).(*response.FindByID)
	}

	var r1 error

	if rf, ok := args.Get(1).(func(context.Context, request.FindByID) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (um *UsecaseMock) TestAddTranslateQuery(ctx context.Context, req request.TestAddTranslateQuery) (*response.TestAddTranslateQuery, error) {
	args := um.Called(ctx, req)

	var r0 *response.TestAddTranslateQuery
	var r1 error

	if rf, ok := args.Get(0).(func(context.Context, request.TestAddTranslateQuery) *response.TestAddTranslateQuery); ok {
		r0 = rf(ctx, req)
	} else if args.Get(0) != nil {
		r0 = args.Get(0).(*response.TestAddTranslateQuery)
	}

	if rf, ok := args.Get(1).(func(context.Context, request.TestAddTranslateQuery) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
