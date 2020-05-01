package user

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
)

// FindByIDRequest to decode get userID from param URL
func FindByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID := chi.URLParam(r, "userID")
	return userReq.FindByID{UserID: userID}, nil
}
