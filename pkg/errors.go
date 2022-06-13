package pkg

import "errors"

var (
	ErrUnsupportedStorageType = errors.New("unsupported storage type")
	ErrNoValue                = errors.New("value is missing")
	ErrInvalidInput           = errors.New("invalid input")
)
