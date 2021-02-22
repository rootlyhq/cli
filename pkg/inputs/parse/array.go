package parse

import "strings"

// Parse an array of strings from an array
func Array(str string) []string {
	var (
		vals    = strings.Split(str, ",")
		trimmed = []string{}
	)
	for _, item := range vals {
		trimmed = append(trimmed, strings.TrimSpace(item))
	}
	if len(trimmed) == 1 && trimmed[0] == "" {
		return []string{}
	}

	return trimmed
}
