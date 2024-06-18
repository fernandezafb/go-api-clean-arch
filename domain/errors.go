package domain

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("entry not found")

// Map of custom errors and their respective status codes used in the custom HTTP error handler
func CustomErrors() map[error]int {

	var errors = make(map[error]int)

	errors[ErrNotFound] = http.StatusNotFound

	return errors
}
