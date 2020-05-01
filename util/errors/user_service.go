package errors

import (
	"net/http"
)

var (
	// UserIDIsRequiredError is missing user id error
	UserIDIsRequiredError = newCustomErr("User ID is required", http.StatusBadRequest)
	// UserNotExistError is user not exist error
	UserNotExistError = newCustomErr("User not exist", http.StatusBadRequest)
	// IncorrectTypeOfUserIDError is incorrect type of user id error
	IncorrectTypeOfUserIDError = newCustomErr("Incorrect type of user ID", http.StatusBadRequest)
)
