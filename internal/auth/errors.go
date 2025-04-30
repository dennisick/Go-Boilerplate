package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenNotFound      = errors.New("token not found")
	ErrInvalidToken       = errors.New("invalid token")
)
