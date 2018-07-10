package airplane

import "errors"

// ErrInValidParam Throw this when invalid param passed to Action
var ErrInValidParam = errors.New("Invalid param passed")

// ErrUnableToExecute Throw this when unable to execute the Action
var ErrUnableToExecute = errors.New("Unable to execute")

// Action Abstraction of Actions
type Action interface {
	Validate() (error, bool)
	Execute() bool
}

// Airplane mode
type Airplane struct {
	Power string
}

//Validate for Airplane mode
func (a Airplane) Validate() (error, bool) {
	if a.Power != "on" && a.Power != "off" {
		return ErrInValidParam, false
	} else {
		return nil, true
	}
}

//Execute the Airplane mode action
func (a Airplane) Execute() bool {
	err, _ := a.Validate()
	if err != nil {
		return false
	}

	w, err := IdentifyDriver()
	if err != nil {
		return false
	}
	w.IdentifyDevice()
	w.Toggle(a.Power)
	return true
}
