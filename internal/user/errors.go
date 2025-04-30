package user

import "errors"

// Custom error types
var (
	ErrUserNotFound = errors.New("user not found")
)
