package errors

import (
	"net/http"
)

//Error Declaration
var (
	InvalidSigningAlgorithm = invalidSigningAlgorithm{}
	InvalidAccessToken      = invalidAccessToken{}
)

// Error signing algorithm
type invalidSigningAlgorithm struct{}

func (invalidSigningAlgorithm) Error() string {
	return "Invalid signing algorithm"
}

func (invalidSigningAlgorithm) StatusCode() int {
	return http.StatusInternalServerError
}

// Error invalid access token
type invalidAccessToken struct{}

func (invalidAccessToken) Error() string {
	return "Invalid access token"
}

func (invalidAccessToken) StatusCode() int {
	return http.StatusBadRequest
}
