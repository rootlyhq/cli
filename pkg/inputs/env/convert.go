package env

import (
	"os"
	"strings"
)

// Get the environmet variable prefix
func GetPrefix() string {
	defaultPrefix := "ROOTLY_"
	if os.Getenv(defaultPrefix+"GH_ACTION") == "true" {
		return "INPUT_"
	}
	return defaultPrefix
}

// Convert a flag name to an environment variable name
func convertName(flagName string) string {
	return GetPrefix() + strings.ReplaceAll(strings.ToUpper(flagName), "-", "_")
}
