package main

import "net/http"

type apiError struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

var (
	ErrUserInvalid    = apiError{Err: "user not valid", Status: http.StatusForbidden}
	ErrMethodInvalid  = apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	ErrInternalServer = apiError{Err: "internal server error", Status: http.StatusInternalServerError}
)

func (e apiError) Error() string {
	return e.Err
}
