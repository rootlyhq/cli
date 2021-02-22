package env

import (
	"os"

	"github.com/rootly-io/cli/pkg/inputs/parse"
)

// Get the value of a string env var
func GetString(name string) string {
	val := os.Getenv(convertName(name))
	return val
}

// Get the value of a string array env var
func GetArray(name string) []string {
	str := os.Getenv(convertName(name))
	return parse.Array(str)
}
