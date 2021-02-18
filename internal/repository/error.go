package repository

import "errors"

// ErrUserNotFound is error when user not found in data storage.
var ErrUserNotFound = errors.New("User not found")

// ErrEmptyParameters is used when was passed empty parameter.
var ErrEmptyParameters = errors.New("Empty parameters")
