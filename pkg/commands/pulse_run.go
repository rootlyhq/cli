package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/rootly-io/cli/pkg/api"
	"github.com/rootly-io/cli/pkg/commands/pulserun"
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/inputs/flags"
	"github.com/rootly-io/cli/pkg/inputs/names"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/cli/pkg/models"
	"github.com/spf13/cobra"
)

var pulseRunCmd = &cobra.Command{
	Use:     "pulse-run",
	Short:   "Run a terminal command and send a pulse with the exit code",
	Example: "rootly pulse-run --api-key \"ABC123\" --summary \"Deploy Website\" --labels \"platform=osx, version=1.12\" sh deploy.sh",
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

		var (
			prog     = args[0]
			progArgs = args[1:]
		)

		apiKey, err := inputs.GetString(names.ApiKeyName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		apiHost, err := inputs.GetString(names.ApiHostName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		summary, err := inputs.GetString(names.PulseSummaryName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}
		if summary == "" {
			summary = prog + " " + strings.Join(progArgs, " ")
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

		log.Success(false, "Got inputs", log.FormatPulse(models.Pulse{
			Summary:        summary,
			Labels:         labels,
			EnvironmentIds: environments,
			ServiceIds:     services,
			StartedAt:      start,
			EndedAt:        time.Now(),
		}))

		client, err := api.GenClient(apiHost)
		if err.Error != nil {
			log.Fatal(err)
		}

		secProvider, err := api.GenSecurityProvider(apiKey)
		if err.Error != nil {
			log.Fatal(err)
		}

		exitCode, err := pulserun.RunProgram(prog, progArgs)
		if err.Error != nil {
			log.Fatal(err)
		}
		labels = append(
			labels,
			map[string]string{"key": "exit_status", "value": fmt.Sprint(exitCode)},
		)
		err = api.CreatePulse(apiHost, models.Pulse{
			Summary:        summary,
			Labels:         labels,
			EnvironmentIds: environments,
			ServiceIds:     services,
			StartedAt:      start,
		}, client, secProvider)
		if err.Error != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pulseRunCmd)

	// Flags
	flags.AddKey(pulseRunCmd)
	flags.AddHost(pulseRunCmd)
	flags.AddPulseLabels(pulseRunCmd)
	flags.AddPulseSummary(pulseRunCmd)
	flags.AddPulseServices(pulseRunCmd)
	flags.AddPulseEnvironments(pulseRunCmd)
	flags.AddOutputDebug(pulseRunCmd)
	flags.AddOutputQuiet(pulseRunCmd)
}
