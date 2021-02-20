package config

import "github.com/spf13/cobra"

// Add a flag for the API key
func AddKeyFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(ApiKeyName),
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info",
	)
}

// Add the standard set of flags specific to pulses
func AddPulseFlags(cmd *cobra.Command) {
	cmd.Flags().StringP(
		string(PulseSummaryName),
		"p",
		"",
		"The summary of the command being ran",
	)
	cmd.Flags().StringArrayP(
		string(PulseLabelsName),
		"l",
		[]string{},
		"Labels associated with the command. Seperated by commands with the key and value of each label being seperated by a colon.",
	)
}
