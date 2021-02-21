package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/rootly-io/cli/pkg/api"
	"github.com/rootly-io/cli/pkg/commands/pulserun"
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

var pulseRunCmd = &cobra.Command{
	Use:     "pulse-run",
	Short:   "Run a terminal command and send a pulse with the exit code",
	Example: "rootly pulse-run --api-key \"ABC123\" --summary \"Deploy Website\" --label=\"Stage=Prod\" sh deploy.sh",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now().UTC()
		log.Info("Getting inputs")

		var (
			prog     = args[0]
			progArgs = args[1:]
		)

		apiKey, err := inputs.GetString(inputs.ApiKeyName, cmd, true)
		if err.Error != nil {
			log.Fatal(err)
		}

		summary, err := inputs.GetString(inputs.PulseSummaryName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}
		if summary == "" {
			summary = prog + " " + strings.Join(progArgs, " ")
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

		exitCode, err := pulserun.RunProgram(prog, progArgs)
		if err.Error != nil {
			log.Fatal(err)
		}
		labels = append(
			labels,
			map[string]string{"key": "Exit Status", "value": fmt.Sprint(exitCode)},
		)

		err = api.CreatePulse(
			api.Pulse{
				Summary:   summary,
				Labels:    labels,
				StartedAt: start,
			},
			client,
			secProvider,
		)
		if err.Error != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pulseRunCmd)

	// Flags
	inputs.AddKeyFlag(pulseRunCmd)
	inputs.AddPulseLabelsFlag(pulseRunCmd)
	inputs.AddPulseSummaryFlag(pulseRunCmd)
}
