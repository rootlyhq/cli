package flags

import (
	"github.com/rootlyhq/cli/pkg/inputs/names"
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
		arrayUsage(name)+" "+mapUsage,
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

// Add a flag for pulse services
func AddPulseSource(cmd *cobra.Command) {
	name := string(names.PulseSourceName)
	cmd.Flags().String(
		name,
		"cli",
		"Source of the pulse",
	)
}

// Add a flag for pulse services
func AddPulseRefs(cmd *cobra.Command) {
	name := string(names.PulseRefsName)
	cmd.Flags().StringP(
		name,
		"r",
		"",
		arrayUsage(name)+" "+mapUsage,
	)
}
