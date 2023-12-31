package utils

import (
	"errors"
	"net/http"
)

// WriteErrorResponse writes an error response to the provided http.ResponseWriter.
func WriteErrorResponse(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	w.Write([]byte(errMsg))
}

// ValidationError creates a validation error response for a specific field.
func ValidationError(field string) error {
	errMsg := "Error Accured :- " + field
	return errors.New(errMsg)
}

func InternalServerError(errMsg string) error {
	return errors.New(errMsg)
}
