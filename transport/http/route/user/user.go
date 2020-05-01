package user

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/clean-architecture/endpoints"
	userDecode "github.com/phungvandat/clean-architecture/transport/http/decode/user"
	"github.com/phungvandat/clean-architecture/transport/http/encode"
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
