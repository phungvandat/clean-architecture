package user

import (
	"context"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	"github.com/phungvandat/clean-architecture/util/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware func
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

// FindByID function handles check input data valid
func (mw validationMiddleware) FindByID(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error) {
	if req.UserID == "" {
		return nil, errors.UserIDIsRequiredError
	}

	if _, err := primitive.ObjectIDFromHex(req.UserID); err != nil {
		return nil, errors.IncorrectTypeOfUserIDError
	}

	return mw.Service.FindByID(ctx, req)
}
