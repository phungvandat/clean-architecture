package repository

import (
	userRepo "github.com/phungvandat/clean-architecture/repository/user"
)

// RepositoryMock struct
type RepositoryMock struct {
	User *userRepo.RepositoryMock
}

// NewRepositoryMock func will return repository mock
func NewRepositoryMock(repoMock RepositoryMock) Repository {
	return Repository{
		User: repoMock.User,
	}
}
