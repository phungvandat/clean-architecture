package repository

import (
	"github.com/phungvandat/clean-architecture/repository/user"
)

// Repository for all service
type Repository struct {
	User user.Repository
}
