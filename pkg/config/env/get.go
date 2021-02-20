package env

import (
	"os"

	"github.com/rootly-io/cli/pkg/log"
)

// Get the value of a string env var
func GetString(name string) string {
	val := os.Getenv(convertName(name))
	return val
}

// Get the value of a string array env var
func GetStringArray(name string) ([]string, log.CtxErr) {
	val := os.Getenv(convertName(name))
	values, err := parseArray(val)
	return values, err
}
