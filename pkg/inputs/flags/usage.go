package flags

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Consistent message telling the user how to use array string flags
func arrayUsage(name string) string {
	return fmt.Sprintf(
		"%v associated with the pulse. Separate each item with a comma.",
		cases.Title(language.English).String(name),
	)
}

// Consistent message telling the user how to use map flags
const mapUsage = "Key-value pair separated with an equal sign (= with no surrounding spaces)"
