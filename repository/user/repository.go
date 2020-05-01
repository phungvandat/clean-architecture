package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
)

// Repository is interface for user repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}
