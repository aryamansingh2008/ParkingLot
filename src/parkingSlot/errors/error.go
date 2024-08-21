package errors

import "errors"

var (
	ErrInvalidID   = errors.New("invalid slot id, id should be in range [1, INT_64_MAX]")
	ErrInvalidPark = errors.New("invalid parking")
)
