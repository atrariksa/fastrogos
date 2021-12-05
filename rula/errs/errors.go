package errs

import (
	"errors"
)

var (
	ErrGeneral         error = errors.New("ErrGeneral")
	ErrInvalidPayload  error = errors.New("ErrInvalidPayload")
	ErrUserNotFound    error = errors.New("ErrUserNotFound")
	ErrInvalidPassword error = errors.New("ErrInvalidPassword")
)
