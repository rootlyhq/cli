package flags

import (
	"github.com/rootly-io/cli/pkg/config"
	"github.com/spf13/cobra"
)

// Add a flag for the API key
func AddKeyFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		config.ApiKeyName,
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info",
	)
}
