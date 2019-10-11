package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/clean-architecture/endpoints"
	"github.com/phungvandat/clean-architecture/transport/http/encode"
	"github.com/phungvandat/clean-architecture/transport/http/options"
	userRoute "github.com/phungvandat/clean-architecture/transport/http/route/user"
	"github.com/rs/cors"
)

func NewHTTPHandler(
	endpoints endpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	r.Use(cors.Handler)

	options := []httpTransport.ServerOption{
		httpTransport.ServerBefore(
			options.LogRequestInfo(logger),
			options.VerifyToken,
		),
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(encode.EncodeError),
	}

	r.Get("/",
		httpTransport.NewServer(
			endpoints.IndexEndpoint,
			httpTransport.NopRequestDecoder,
			httpTransport.EncodeJSONResponse,
			options...,
		).ServeHTTP)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", userRoute.UserRoute(endpoints, options))
	})

	return r
}
