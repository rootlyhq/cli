package commands

import (
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

// The main root command for rootly
var rootCmd = &cobra.Command{
	Use:   "rootly",
	Short: "Command-line Tool for rootly.io",
	Long: `
					_   _
	_ __ ___   ___ | |_| |_   _
	| '__/ _ \ / _ \| __| | | | |
	| | | (_) | (_) | |_| | |_| |
	|_|  \___/ \___/ \__|_|\__, |
							|___/

	Command-line Tool for rootly.io`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Welcome to the void 😎")
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(log.CtxErr{
			Context: "Failed to execute root command",
			Error:   err,
		})
	}
}
