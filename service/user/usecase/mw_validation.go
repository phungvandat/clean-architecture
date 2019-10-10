package usecase

import (
	"context"

	"github.com/phungvandat/identity-service/model/request"
	"github.com/phungvandat/identity-service/model/response"
	"github.com/phungvandat/identity-service/service/user"
	"github.com/phungvandat/identity-service/util/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type validationMiddleware struct {
	user.Service
}

func ValidationMiddleware() func(user.Service) user.Service {
	return func(next user.Service) user.Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) FindByID(ctx context.Context, req request.FindByID) (*response.FindByID, error) {
	if req.UserID == "" {
		return nil, errors.UserIDIsRequiredError
	}

	if _, err := primitive.ObjectIDFromHex(req.UserID); err != nil {
		return nil, errors.IncorrectTypeOfUserIDError
	}

	return mw.Service.FindByID(ctx, req)
}

func (mw validationMiddleware) TestAddTranslateQuery(ctx context.Context, req request.TestAddTranslateQuery) (*response.TestAddTranslateQuery, error) {
	return mw.Service.TestAddTranslateQuery(ctx, req)
}
