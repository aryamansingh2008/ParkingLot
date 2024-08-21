package errors

import "errors"

var (
	ErrInvalidType         = errors.New("invalid vehicle type")
	ErrInvalidRegistration = errors.New("invalid vehicle registration number")
)
