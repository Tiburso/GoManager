package common

import "errors"

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrPermissionDenied = errors.New("permission denied")
	ErrAlreadyExist     = errors.New("resource already exists")
	ErrNotExist         = errors.New("resource does not exist")
)
