package user

import (
	"github.com/go-kit/kit/endpoint"
)

type UserEndpoint struct {
	FindByID              endpoint.Endpoint
	TestAddTranslateQuery endpoint.Endpoint
}
