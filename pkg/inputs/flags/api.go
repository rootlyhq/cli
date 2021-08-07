package flags

import (
	"github.com/rootlyhq/cli/pkg/inputs/names"
	"github.com/spf13/cobra"
)

// Add a flag for the API key
func AddKey(cmd *cobra.Command) {
	cmd.Flags().StringP(
		string(names.ApiKeyName),
		"k",
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info.",
	)
}

// Add a flag for the API host
func AddHost(cmd *cobra.Command) {
	cmd.Flags().String(
		string(names.ApiHostName),
		"https://api.rootly.io",
		"Host URL for the rootly api.",
	)
}
