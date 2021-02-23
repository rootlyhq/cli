package commands

import (
	"os"
	"strings"
	"time"

	"github.com/rootly-io/cli/pkg/api"
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/inputs/env"
	"github.com/rootly-io/cli/pkg/inputs/flags"
	"github.com/rootly-io/cli/pkg/inputs/names"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/cli/pkg/models"
	"github.com/spf13/cobra"
)

var pulseCmd = &cobra.Command{
	Use:     "pulse",
	Short:   "Send a pulse",
	Example: "rootly pulse --api-key \"ABC123\" --labels \"platform=osx, version=1.12\" Deployed Site",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now().UTC()

		quiet, err := inputs.GetBool(names.OutputQuietName, cmd)
		if err.Error != nil {
			log.Fatal(err)
		}
		log.Quiet = quiet

		debug, err := inputs.GetBool(names.OutputDebugName, cmd)
		if err.Error != nil {
			log.Fatal(err)
		}
		log.Debug = debug

		log.Info("Getting pulse inputs")

		apiKey, err := inputs.GetString(names.ApiKeyName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		apiHost, err := inputs.GetString(names.ApiHostName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		summary := strings.Join(args, " ")
		if summary == "" {
			summary = os.Getenv(env.GetPrefix() + "SUMMARY")
			if summary == "" {
				log.Fatal(log.NewErr("No summary provided"))
			}
		}

		labels, err := inputs.GetStringSimpleMapArray(names.PulseLabelsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		environments, err := inputs.GetArray(names.PulseEnvironmentsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		services, err := inputs.GetArray(names.PulseServicesName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		pulse := models.Pulse{
			Summary:        summary,
			Labels:         labels,
			EnvironmentIds: environments,
			ServiceIds:     services,
			StartedAt:      start,
			EndedAt:        time.Now().UTC(),
		}
		log.Success(false, "Got pulse inputs", log.FormatPulse(pulse))

		client, err := api.GenClient(apiHost)
		if err.Error != nil {
			log.Fatal(err)
		}

		secProvider, err := api.GenSecurityProvider(apiKey)
		if err.Error != nil {
			log.Fatal(err)
		}

		err = api.CreatePulse(apiHost, pulse, client, secProvider)
		if err.Error != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pulseCmd)

	// Flags
	flags.AddKey(pulseCmd)
	flags.AddHost(pulseCmd)
	flags.AddPulseLabels(pulseCmd)
	flags.AddPulseServices(pulseCmd)
	flags.AddPulseEnvironments(pulseCmd)
	flags.AddOutputDebug(pulseCmd)
	flags.AddOutputQuiet(pulseCmd)
}
