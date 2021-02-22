package env

import (
	"os"
	"strings"
)

// Convert a flag name to an environment variable name
func convertName(flagName string) string {
	prefix := "ROOTLY_CLI_"
	if os.Getenv(prefix+"GH_ACTION") == "true" {
		prefix = "INPUT_"
	}
	return prefix + strings.ReplaceAll(strings.ToUpper(flagName), "-", "_")
}
