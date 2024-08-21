package errors

import "errors"

var (
	SlotIdNotFound = errors.New("could not find a slot with given slot ID")
)
