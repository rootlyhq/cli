package flags

import (
	"fmt"
	"strings"
)

// Consistent message telling the user how to use array string flags
func arrayUsage(name string) string {
	return fmt.Sprintf(
		"%v associated with the pulse. Separate each item with a comma.",
		strings.Title(name),
	)
}
