package common

import "errors"

var (
	ErrInvalidParam = errors.New("invalid parameter")
	ErrServer       = errors.New("server error")
	ErrUnauthorized = errors.New("unauthorized")
)
