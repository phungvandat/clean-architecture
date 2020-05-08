package errors

import "net/http"

var (
	// OccurredError error
	OccurredError = newCustomErr("An error occurred, Please try again", http.StatusInternalServerError)
)
