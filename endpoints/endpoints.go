package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/clean-architecture/endpoints/index"
	"github.com/phungvandat/clean-architecture/endpoints/user"
	"github.com/phungvandat/clean-architecture/service"
)

type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
	UserEndpoint  user.UserEndpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		IndexEndpoint: index.MakeIndexEndpoints(),
		UserEndpoint: user.UserEndpoint{
			FindByID:              user.MakeFindByIDEndpoint(s),
			TestAddTranslateQuery: user.MakeTestAddTranslateQuery(s),
		},
	}
}
