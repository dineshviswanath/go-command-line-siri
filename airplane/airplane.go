package airplane

import "github.com/go-commandline-siri/util"

// Airplane mode
type Airplane struct {
	Power string
}

//Validate for Airplane mode
func (a Airplane) Validate() (error, bool) {
	if a.Power != "on" && a.Power != "off" {
		return util.ErrInValidParam, false
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
