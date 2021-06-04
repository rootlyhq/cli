package log

import "github.com/gleich/lumber"

var (
	// If no logs should be outputted
	Quiet = false
	// If all logs should be outputted
	Debug = false
)

// Log a given error fatally to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Fatal(err CtxErr) {
	lumber.Fatal(err.Error, err.Context)
}

// Log some info to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Info(ctx ...interface{}) {
	if Debug {
		lumber.Info(ctx...)
	}
}

// Log a success to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Success(forceOut bool, ctx ...interface{}) {
	if (forceOut || Debug) && !Quiet {
		lumber.Success(ctx...)
	}
}

// Log a warning to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Warning(ctx ...interface{}) {
	if Debug {
		lumber.Warning(ctx...)
	}
}
