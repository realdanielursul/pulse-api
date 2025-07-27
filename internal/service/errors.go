package service

import "errors"

var (
	ErrInvalidRegion   = errors.New("invalid region")
	ErrCountryNotFound = errors.New("country not found")
)
