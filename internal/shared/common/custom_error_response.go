package common

import "errors"

var (
	ErrInvalidInput      = errors.New("Invalid input")
	ErrNotFound          = errors.New("Not found")
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrInternalError     = errors.New("Internal server error")
	ErrInvalidPassword   = errors.New("Password must be at least 8 characters, contains uppercase, lowercase, number, and special characters")
	ErrCoflict           = errors.New("Email already registered")
	ErrWrongPassword     = errors.New("Invalid email or password")
	ErrStockAvailability = errors.New("The quantity of the product exceeds the available stock")
)
