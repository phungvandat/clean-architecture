package user

import (
	"context"
	"net/http"
	"time"

	"strconv"

	"github.com/go-chi/chi"
	"github.com/phungvandat/clean-architecture/model/request"
)

// FindByIDRequest to decode get userID from param URL
func FindByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID := chi.URLParam(r, "userID")
	return request.FindByID{UserID: userID}, nil
}

// TestAddTranslateQueryRequest to decode for test add translate query request
func TestAddTranslateQueryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	timeNum, err := strconv.ParseInt(r.URL.Query().Get("createdAt"), 10, 64)

	var createdAt *time.Time
	if err == nil {
		// Client send milisecond
		crt := time.Unix(timeNum/1000, 0)
		createdAt = &crt
	}

	availableParse, err := strconv.ParseBool(r.URL.Query().Get("available"))
	var available *bool
	if err == nil {
		available = &availableParse
	}

	return request.TestAddTranslateQuery{
		CreatedAt: createdAt,
		Fullname:  r.URL.Query().Get("fullname"),
		Available: available,
	}, nil
}
