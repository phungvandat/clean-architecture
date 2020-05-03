package user

import (
	"github.com/go-kit/kit/endpoint"
)

// UserEndpoint struct contain all endpoint of user service
type UserEndpoint struct {
	FindByID endpoint.Endpoint
	Find     endpoint.Endpoint
}
