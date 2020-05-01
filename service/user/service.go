package user

import (
	"context"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
)

// Service for user
type Service interface {
	FindByID(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error)
}
