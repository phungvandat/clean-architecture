package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
)

// FindByIDRequest to decode get userID from param URL
func FindByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID := chi.URLParam(r, "userID")
	return userReq.FindByID{UserID: userID}, nil
}

// FindRequest to decode data
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	reqBody := userReq.Find{}

	qValues := r.URL.Query()

	if len(qValues["fullname"]) > 0 {
		reqBody.Fullname = qValues["fullname"][0]
	}

	return reqBody, nil
}

// CreateRequest to decode data
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	reqBody := userReq.Create{}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return nil, err
	}

	return reqBody, nil
}
