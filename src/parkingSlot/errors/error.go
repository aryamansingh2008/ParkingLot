package errors

import "errors"

var (
	InvalidID                 = errors.New("invalid slot id, id should be in range [1, INT_64_MAX]")
	InvalidParkOccupied       = errors.New("invalid parking, slot is already occupied")
	InvalidParkCantAccomodate = errors.New("invalid parking, can't accomodate")
)
