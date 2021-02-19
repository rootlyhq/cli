package log

// A error with context that the user can understand
type CtxErr struct {
	Error   error
	Context string
}
