package model

import "errors"

var (
	ErrUserDoesNotExist = errors.New("requested user does not exist")
	ErrNilUserID = errors.New("invalid/nil user id")
)
