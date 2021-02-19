package log

import "github.com/Matt-Gleich/lumber"

// Log a given error to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Error(err CtxErr) {
	lumber.Fatal(err.Error, err.Context)
}

// Log some info to the console
// At the moment this function is a simple wrapper around
// the lumber module but exists for easy switching to other
// log formats
func Info(ctx ...interface{}) {
	lumber.Info(ctx...)
}
