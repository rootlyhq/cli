package log

import (
	"fmt"
	"strings"
	"time"

	"github.com/rootlyhq/cli/pkg/models"
)

// Format a pulse for outputting
func FormatPulse(pulse models.Pulse) string {
	delimiter := ", "

	return strings.TrimSuffix(fmt.Sprintf(`

  Summary:      %v
  Started At:   %v
  Ended At      %v
  Services:     %v
  Environments: %v
  Labels:       %v
  Source:       %v
  Refs:         %v`,
		emptyReplace(pulse.Summary),
		emptyReplace(pulse.StartedAt.Format(time.RFC822)),
		emptyReplace(pulse.EndedAt.Format(time.RFC822)),
		emptyReplace(strings.Join(pulse.ServiceIds, delimiter)),
		emptyReplace(strings.Join(pulse.EnvironmentIds, delimiter)),
		emptyReplace(formatMaps(pulse.Labels)),
		emptyReplace(pulse.Source),
		emptyReplace(formatMaps(pulse.Refs)),
	), "\n")
}

// Format a map
func formatMaps(maps []map[string]string) string {
	var fmtMap string
	for i, mapData := range maps {
		prefix := "    "
		if i == 0 {
			prefix = "\n    "
		}
		fmtMap = fmtMap + fmt.Sprintf(
			"%v%v = %v\n",
			prefix,
			mapData["key"],
			mapData["value"],
		)
	}
	return fmtMap
}

// Replace empty string values with None
func emptyReplace(str string) string {
	if str == "" {
		return "None"
	}
	return str
}
