package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/request"
	"github.com/phungvandat/clean-architecture/model/response"
)

type Service interface {
	FindByID(ctx context.Context, req request.FindByID) (*response.FindByID, error)
	TestAddTranslateQuery(ctx context.Context, req request.TestAddTranslateQuery) (*response.TestAddTranslateQuery, error)
}
