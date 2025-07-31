package service

import "errors"

var (
	ErrAccessDenied           = errors.New("access denied")
	ErrCannotParseToken       = errors.New("cannot parse token")
	ErrCannotSignToken        = errors.New("cannot sign token")
	ErrCountryNotFound        = errors.New("country not found")
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrInvalidRegion          = errors.New("invalid region")
	ErrLoginAlreadyExists     = errors.New("login already exists")
	ErrPostNotFound           = errors.New("post not found")
	ErrPhoneAlreadyExists     = errors.New("phone already exists")
	ErrUserNotFound           = errors.New("user not found")
)
