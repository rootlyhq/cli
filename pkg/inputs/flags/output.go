package flags

import (
	"github.com/rootly-io/cli/pkg/inputs/names"
	"github.com/spf13/cobra"
)

// Add output silent flag
func AddOutputSilent(cmd *cobra.Command) {
	cmd.Flags().Bool(
		string(names.OutputSilentName),
		false,
		"If the output should be silent",
	)
}

// Add output debug flag
func AddOutputDebug(cmd *cobra.Command) {
	cmd.Flags().Bool(
		string(names.OutputDebugName),
		false,
		"Provide more output than usual for debugging",
	)
}
