package errors

import "errors"

var (
	InvalidID   = errors.New("invalid slot id, id should be in range [1, INT_64_MAX]")
	InvalidPark = errors.New("invalid parking")
)
