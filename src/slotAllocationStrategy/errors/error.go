package errors

import "errors"

var (
	NoValidSlotFound = errors.New("could not find a valid slot for the vehicle")
)
