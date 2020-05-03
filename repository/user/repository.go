package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
)

// Repository is interface for user repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Find(ctx context.Context, conditions userInput.FindConditions) ([]*domain.User, error)
}
