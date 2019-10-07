package user

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phungvandat/identity-service/model/request"
)

func FindByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID := chi.URLParam(r, "userID")
	return request.FindByID{UserID: userID}, nil
}
