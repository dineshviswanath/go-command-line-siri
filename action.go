package main

// Abstraction of Actions
type Action interface {
	Validate() (error, bool)
	Execute() bool
}
