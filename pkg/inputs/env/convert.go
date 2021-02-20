package env

import "strings"

// Convert a flag name to an environment variable name
func convertName(flagName string) string {
	return "ROOTLY_CLI_" + strings.ReplaceAll(strings.ToUpper(flagName), "-", "_")
}
