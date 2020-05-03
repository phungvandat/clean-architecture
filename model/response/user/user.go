package user

import (
	"github.com/phungvandat/clean-architecture/model/domain"
)

// FindByID struct
type FindByID struct {
	User *domain.User `json:"user"`
}

// Find struct
type Find struct {
	Users []*domain.User `json:"users"`
}
