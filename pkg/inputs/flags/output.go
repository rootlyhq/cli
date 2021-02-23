package flags

import (
	"github.com/rootly-io/cli/pkg/inputs/names"
	"github.com/spf13/cobra"
)

// Add output quiet flag
func AddOutputQuiet(cmd *cobra.Command) {
	cmd.Flags().BoolP(
		string(names.OutputQuietName),
		"q",
		false,
		"If the output should be quiet",
	)
}

// Add output debug flag
func AddOutputDebug(cmd *cobra.Command) {
	cmd.Flags().BoolP(
		string(names.OutputDebugName),
		"d",
		false,
		"Provide more output than usual for debugging",
	)
}
