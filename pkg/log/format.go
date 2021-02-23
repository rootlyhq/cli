package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/rootly-io/cli/pkg/models"
)

// Format a pulse for outputting
func FormatPulse(pulse models.Pulse) string {
	delimiter := ", "

	// Putting labels in a readable format
	var fmtLabels string
	for i, label := range pulse.Labels {
		prefix := "    "
		if i == 0 {
			prefix = "\n    "
		}
		fmtLabels = fmtLabels + fmt.Sprintf(
			"%v%v = %v\n",
			prefix,
			label["key"],
			label["value"],
		)
	}

	return strings.TrimSuffix(fmt.Sprintf(`

  Summary:      %v
  Started At:   %v
  Ended At      %v
  Services:     %v
  Environments: %v
  Labels:       %v`,
		emptyReplace(pulse.Summary),
		emptyReplace(pulse.StartedAt.Format(time.RFC822)),
		emptyReplace(pulse.EndedAt.Format(time.RFC822)),
		emptyReplace(strings.Join(pulse.ServiceIds, delimiter)),
		emptyReplace(strings.Join(pulse.EnvironmentIds, delimiter)),
		emptyReplace(fmtLabels),
	), "\n")
}

// Replace empty string values with None
func emptyReplace(str string) string {
	if str == "" {
		return "None"
	}
	return str
}
