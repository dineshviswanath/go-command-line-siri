package core

import "errors"

// ErrInValidParam Throw this when invalid param passed to Action
var ErrInValidParam = errors.New("Invalid param passed")

// ErrUnableToExecute Throw this when unable to execute the Action
var ErrUnableToExecute = errors.New("Unable to execute")

// Action Abstraction of Actions
type Action interface {
	Validate() error
	Execute() error
}
