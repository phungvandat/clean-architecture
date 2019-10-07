package index

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Data response client
type IndexData struct {
	Title string `json:"title"`
}

func MakeIndexEndpoints() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return IndexData{
			Title: "Identity service",
		}, nil
	}
}
