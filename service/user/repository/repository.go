package repository

import (
	"context"

	"github.com/phungvandat/identity-service/model/domain"
)

// Repository is interface for user repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}
