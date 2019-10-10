package service

import (
	"github.com/phungvandat/clean-architecture/service/user"
)

type Service struct {
	UserService user.Service
}
