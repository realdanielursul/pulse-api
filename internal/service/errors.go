package service

import "errors"

var (
	ErrInvalidRegion   = errors.New("invalid region")
	ErrCountryNotFound = errors.New("country not found")
)

var (
	ErrLoginAlreadyExists = errors.New("login already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPhoneAlreadyExists = errors.New("phone already exists")
)
