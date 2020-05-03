package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/clean-architecture/endpoints/index"
	"github.com/phungvandat/clean-architecture/endpoints/user"
	"github.com/phungvandat/clean-architecture/service"
)

// Endpoints struct
type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
	UserEndpoint  user.UserEndpoint
}

// MakeServerEndpoints function return summary of all endpoint
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		IndexEndpoint: index.MakeIndexEndpoints(),
		UserEndpoint: user.UserEndpoint{
			FindByID: user.MakeFindByIDEndpoint(s),
			Find:     user.MakeFindEndpoint(s),
		},
	}
}
