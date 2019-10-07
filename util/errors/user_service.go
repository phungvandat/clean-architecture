package errors

import (
	"net/http"
)

var (
	UserIDIsRequiredError      = userIDIsRequiredError{}
	UserNotExistError          = userNotExistError{}
	IncorrectTypeOfUserIDError = incorrectTypeOfUserIDError{}
)

type userIDIsRequiredError struct{}

func (userIDIsRequiredError) Error() string {
	return "User ID is required"
}

func (userIDIsRequiredError) StatusCode() int {
	return http.StatusBadRequest
}

type userNotExistError struct{}

func (userNotExistError) Error() string {
	return "User not exist"
}

func (userNotExistError) StatusCode() int {
	return http.StatusBadRequest
}

type incorrectTypeOfUserIDError struct{}

func (incorrectTypeOfUserIDError) Error() string {
	return "Incorrect type of user ID"
}

func (incorrectTypeOfUserIDError) StatusCode() int {
	return http.StatusBadRequest
}
