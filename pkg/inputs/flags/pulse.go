package flags

import (
	"github.com/rootly-io/cli/pkg/inputs/names"
	"github.com/spf13/cobra"
)

// Add a flag for a pulse summary
func AddPulseSummary(cmd *cobra.Command) {
	cmd.Flags().String(
		string(names.PulseSummaryName),
		"",
		"Summary of the pulse",
	)
}

// Add a flag for pulse labels
func AddPulseLabels(cmd *cobra.Command) {
	name := string(names.PulseLabelsName)
	cmd.Flags().StringP(
		name,
		"l",
		"",
		arrayUsage(name)+" Key-value pair separated with equal sign (=)",
	)
}

// Add a flag for pulse services
func AddPulseServices(cmd *cobra.Command) {
	name := string(names.PulseServicesName)
	cmd.Flags().StringP(
		name,
		"s",
		"",
		arrayUsage(name),
	)
}

// Add a flag for pulse services
func AddPulseEnvironments(cmd *cobra.Command) {
	name := string(names.PulseEnvironmentsName)
	cmd.Flags().StringP(
		name,
		"e",
		"",
		arrayUsage(name),
	)
}
