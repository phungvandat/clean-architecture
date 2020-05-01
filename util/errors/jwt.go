package errors

import (
	"net/http"
)

// Error Declaration
var (
	// Error signing algorithm
	InvalidSigningAlgorithm = newCustomErr("Invalid signing algorithm", http.StatusInternalServerError)
	// Error invalid access token
	InvalidAccessToken = newCustomErr("Invalid access token", http.StatusBadRequest)
)
