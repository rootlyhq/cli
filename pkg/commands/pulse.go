package commands

import (
	"strings"
	"time"

	"github.com/rootly-io/cli/pkg/api"
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/cli/pkg/models"
	"github.com/spf13/cobra"
)

var pulseCmd = &cobra.Command{
	Use:     "pulse",
	Short:   "Send a pulse",
	Example: "rootly pulse --api-key \"ABC123\" --label \"Version=3\" --label \"Deployed By=Harry Potter\" Deployed Site",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now().UTC()
		log.Info("Getting pulse inputs")

		apiKey, err := inputs.GetString(inputs.ApiKeyName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		summary := strings.Join(args, " ")

		labels, err := inputs.GetStringSimpleMapArray(inputs.PulseLabelsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		environments, err := inputs.GetStringArray(inputs.PulseEnvironmentsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		services, err := inputs.GetStringArray(inputs.PulseServicesName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		pulse := models.Pulse{
			Summary:        summary,
			Labels:         labels,
			EnvironmentIds: environments,
			ServiceIds:     services,
			StartedAt:      start,
		}
		log.Success("Got pulse inputs", log.FormatPulse(pulse))

		client, err := api.GenClient()
		if err.Error != nil {
			log.Fatal(err)
		}

		secProvider, err := api.GenSecurityProvider(apiKey)
		if err.Error != nil {
			log.Fatal(err)
		}

		err = api.CreatePulse(pulse, client, secProvider)
		if err.Error != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pulseCmd)

	// Flags
	inputs.AddKeyFlag(pulseCmd)
	inputs.AddPulseLabelsFlag(pulseCmd)
	inputs.AddPulseServicesFlag(pulseCmd)
	inputs.AddPulseEnvironmentsFlag(pulseCmd)
}
