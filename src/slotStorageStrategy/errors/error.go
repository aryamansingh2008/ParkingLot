package errors

import "errors"

var (
	ErrSlotIdNotFound = errors.New("could not find a slot with given slot ID")
)
