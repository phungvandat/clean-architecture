package options

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
)

func LogRequestInfo(logger log.Logger) func(ctx context.Context, req *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		logger.Log("Method", req.Method, "Route", req.RequestURI)
		return ctx
	}
}
