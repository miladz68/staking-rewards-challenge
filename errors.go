package main

import "errors"

var (
	ErrUnimplemented   = errors.New("Unimplemented")
	ErrInvalidArgument = errors.New("Invalid Argument")
	ErrNotFound        = errors.New("Not Found")
)
