package service

import "errors"

var (
	ErrInvalidRegion   = errors.New("invalid region")
	ErrCountryNotFound = errors.New("country not found")
)

var (
	ErrLoginAlreadyExists     = errors.New("login already exists")
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrPhoneAlreadyExists     = errors.New("phone already exists")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")

	ErrCannotSignToken  = errors.New("cannot sign token")
	ErrCannotParseToken = errors.New("cannot parse token")

	ErrAccessDenied = errors.New("access denied")
)
