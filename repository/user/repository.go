package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
)

// Repository is interface for user repository
type Repository interface {
	FindByID(ctx context.Context, id string, options ...*repository.RepoOptions) (*domain.User, error)
	Find(ctx context.Context, conditions userInput.FindConditions, options ...*repository.RepoOptions) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User, options ...*repository.RepoOptions) (*domain.User, error)
}
