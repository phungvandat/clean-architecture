package response

import (
	"github.com/phungvandat/identity-service/model/domain"
)

type FindByID struct {
	User *domain.User `json:"user"`
}

type TestAddTranslateQuery struct {
	Users []*domain.User `json:"users"`
}
