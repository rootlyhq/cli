package commands

import (
	"github.com/rootly-io/cli/pkg/config"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

var pulseCmd = &cobra.Command{
	Use:     "pulse",
	Short:   "Send a pulse",
	Example: "rootly pulse --summary=\"Deployed Site\" --api-key=\"ABC123\" --labels=\"Version|3,Deployed By|Harry Potter\"",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Sent a pulse")
	},
}

func init() {
	rootCmd.AddCommand(pulseCmd)

	// Flags
	config.AddKeyFlag(pulseCmd)
	config.AddPulseFlags(pulseCmd)
}
