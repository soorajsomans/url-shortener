package errors

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrInvalidURL  = errors.New("invalid url")
	ErrURLExpired  = errors.New("url expired")
	ErrAliasExists = errors.New("alias already exists")
)
