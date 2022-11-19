package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("NOT FOUND")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("ALREADY EXIST")
	// ErrEmailAlreadyExist will throw if the given request-body or params is not valid
	ErrEmailAlreadyExist = errors.New("EMAIL ALREADY EXIST")
	// ErrUnauthorized will throw if the token is invalid
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)
