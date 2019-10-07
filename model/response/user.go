package response

import (
	"github.com/phungvandat/identity-service/model/domain"
)

type FindByID struct{
	User *domain.User `json:"user"`
}