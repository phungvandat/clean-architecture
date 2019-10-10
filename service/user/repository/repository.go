package repository

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/util/engine"
)

// Repository is interface for user repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	TestAddTranslateQuery(ctx context.Context, query *engine.Query) ([]*domain.User, error)
}
