package inputs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Add a flag for the API key
func AddKeyFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(ApiKeyName),
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info.",
	)
}

// Add a flag for the API host
func AddHostFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(ApiHostName),
		"https://api.rootly.io",
		"Host URL for the rootly api.",
	)
}

// Add a flag for a pulse summary
func AddPulseSummaryFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(PulseSummaryName),
		"",
		"Summary of the pulse",
	)
}

// Consistent message telling the user how to use array string flags
func arrayUsage(name string) string {
	return fmt.Sprintf(
		"%v associated with the pulse. Separate each item with a comma.",
		strings.Title(name),
	)
}

// Add a flag for pulse labels
func AddPulseLabelsFlag(cmd *cobra.Command) {
	name := string(PulseLabelsName)
	cmd.Flags().String(
		name,
		"",
		arrayUsage(name)+" Key-value pair seperated with equal sign (=)",
	)
}

// Add a flag for pulse services
func AddPulseServicesFlag(cmd *cobra.Command) {
	name := string(PulseServicesName)
	cmd.Flags().String(
		name,
		"",
		arrayUsage(name),
	)
}

// Add a flag for pulse services
func AddPulseEnvironmentsFlag(cmd *cobra.Command) {
	name := string(PulseEnvironmentsName)
	cmd.Flags().String(
		name,
		"",
		arrayUsage(name),
	)
}
