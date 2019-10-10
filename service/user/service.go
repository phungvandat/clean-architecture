package user

import (
	"context"

	"github.com/phungvandat/identity-service/model/request"
	"github.com/phungvandat/identity-service/model/response"
)

type Service interface {
	FindByID(ctx context.Context, req request.FindByID) (*response.FindByID, error)
	TestAddTranslateQuery(ctx context.Context, req request.TestAddTranslateQuery) (*response.TestAddTranslateQuery, error)
}
