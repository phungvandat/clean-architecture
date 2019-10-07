package user

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/identity-service/endpoints"
	userDecode "github.com/phungvandat/identity-service/http/decode/user"
	"github.com/phungvandat/identity-service/http/encode"
)

// UserRoute route
func UserRoute(
	endpoints endpoints.Endpoints,
	options []httptransport.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		// Find by ID
		r.Get("/{userID}", httptransport.NewServer(
			endpoints.UserEndpoint.FindByID,
			userDecode.FindByIDRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
