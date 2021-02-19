package log

import "errors"

// Create a brand new error
func NewErr(context string) CtxErr {
	return CtxErr{
		Context: context,
		Error:   errors.New(context),
	}
}

// A error with context that the user can understand
type CtxErr struct {
	Error   error
	Context string
}
