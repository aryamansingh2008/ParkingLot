package errors

import "errors"

var (
	ErrNoValidSlotFound = errors.New("could not find a valid slot for the vehicle")
)
