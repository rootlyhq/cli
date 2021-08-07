package env

import (
	"os"

	"github.com/rootlyhq/cli/pkg/inputs/parse"
	"github.com/rootlyhq/cli/pkg/models"
)

// Get the value of a string env var
func GetString(name models.ConfigPiece) string {
	return os.Getenv(convertName(string(name)))
}

// Get the value of a string array env var
func GetArray(name models.ConfigPiece) []string {
	str := os.Getenv(convertName(string(name)))
	return parse.Array(str)
}
