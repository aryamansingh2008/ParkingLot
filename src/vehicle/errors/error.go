package errors

import "errors"

var (
	InvalidType         = errors.New("invalid vehicle type")
	InvalidRegistration = errors.New("invalid vehicle registration number")
)
