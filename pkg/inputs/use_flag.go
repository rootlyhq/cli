package inputs

import "github.com/spf13/cobra"

// Add a flag for the API key
func AddKeyFlag(cmd *cobra.Command) {
	cmd.Flags().String(
		string(ApiKeyName),
		"",
		"api key generated from rootly.io. See https://rootly.io/api#section/How-to-generate-an-API-Key for more info",
	)
}

// Add a flag for a summary
func AddPulseSummaryFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(
		string(PulseSummaryName),
		"p",
		"",
		"The summary of the command being ran",
	)
}

// Add a flag for labels
func AddPulseLabelsFlag(cmd *cobra.Command) {
	cmd.Flags().StringArrayP(
		string(PulseLabelsName),
		"l",
		[]string{},
		"Labels associated with the command. Give a new flag for each label and separate the key from the value with an equals sign.",
	)
}
