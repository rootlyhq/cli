package commands

import (
	"github.com/rootly-io/cli/pkg/inputs"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

var pulseRunCmd = &cobra.Command{
	Use:     "pulse-run",
	Short:   "Run a command and send a pulse with the exit code",
	Example: "rootly pulse-run --api-key\"ABC123\" --summary\"Deploy Website\" --label=\"Stage|#|Prod\" sh deploy.sh",
	Run: func(cmd *cobra.Command, args []string) {
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
			summary = prog
		}

		labels, err := inputs.GetStringSimpleMapArray(inputs.PulseLabelsName, cmd, false)
		if err.Error != nil {
			log.Fatal(err)
		}

		log.Success("Got inputs")

	},
}

func init() {
	rootCmd.AddCommand(pulseRunCmd)

	// Flags
	inputs.AddKeyFlag(pulseRunCmd)
	inputs.AddPulseFlags(pulseRunCmd)
}
