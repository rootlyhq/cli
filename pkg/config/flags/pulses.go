package flags

import (
	"github.com/rootly-io/cli/pkg/config"
	"github.com/spf13/cobra"
)

// Add the standard set of flags specific to pulses
func AddPulseFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(
		config.PulseSummaryName,
		"p",
		"",
		"The summary of the command being ran",
	)
	cmd.Flags().StringArrayP(
		config.PulseLabelsName,
		"l",
		[]string{},
		"Labels associated with the command. Seperated by commands with the key and value of each label being seperated by a colon.",
	)
}
