package service

import (
	"errors"
)

var (
	// ErrShopNotFound is returned when a shop is not found.
	ErrShopNotFound = errors.New("shop not found")

	// ErrShopAlreadyExists is returned when a shop with the same name already exists.
	ErrShopAlreadyExists = errors.New("shop already exists")

	// Add more custom error variables as needed for your specific application.
)
