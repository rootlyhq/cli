package commands

import (
	"github.com/rootly-io/cli/pkg/api"
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

var pulseCmd = &cobra.Command{
	Use:     "pulse",
	Short:   "Send a pulse",
	Example: "rootly pulse --summary=\"Deployed Site\" --api-key=\"ABC123\" --label=\"Version|#|3\" --label=\"Deployed By|#|Harry Potter\"",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Getting inputs")

		apiKey, err := inputs.GetString(inputs.ApiKeyName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		summary, err := inputs.GetString(inputs.PulseSummaryName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		labels, err := inputs.GetStringSimpleMapArray(inputs.PulseLabelsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		log.Success("Got inputs")

		client, err := api.GenClient()
		if err.Error != nil {
			log.Fatal(err)
		}

		secProvider, err := api.GenSecurityProvider(apiKey)
		if err.Error != nil {
			log.Fatal(err)
		}

		err = api.CreatePulse(api.Pulse{Summary: summary, Labels: labels}, client, secProvider)
		if err.Error != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pulseCmd)

	// Flags
	inputs.AddKeyFlag(pulseCmd)
	inputs.AddPulseFlags(pulseCmd)
}