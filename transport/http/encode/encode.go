package encode

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	kithttp "github.com/go-kit/kit/transport/http"
)

// encodeResponse is the common method to encode all response types to the client.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}

// EncodeFileResponse func
func EncodeFileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "image/jpeg")
	file := response.(*os.File)
	defer file.Close()
	_, err := io.Copy(w, file)
	return err
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	// maybe we can be smart here by returning text/json error based on request's
	// content-type header
	encodeJSONError(ctx, err, w)
}

func encodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// custom headers
	if headerer, ok := err.(kithttp.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	// custome code
	if sc, ok := err.(kithttp.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	// enforce json response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": err.Error(),
	})
}
