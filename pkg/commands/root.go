package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/gleich/release"
	"github.com/rootlyhq/cli/pkg/log"
	"github.com/spf13/cobra"
)

// The main root command for rootly
var rootCmd = &cobra.Command{
	Use:   "rootly",
	Short: "Command-line Tool for rootly.io",
	Long: `  	           _   _
   _ __ ___   ___ | |_| |_   _
  | '__/ _ \ / _ \| __| | | | |
  | | | (_) | (_) | |_| | |_| |
  |_|  \___/ \___/ \__|_|\__, |
  			  |___/

  Command-line Tool for rootly.io`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(log.CtxErr{
				Context: "Failed to get version command",
			})
		}

		if versionFlag {
			version := "v1.0.1"
			outdated, v, err := release.Check(version, "https://github.com/rootlyhq/cli")

			if err != nil {
				log.Warning("Failed to get latest version of rootly")
			}
			if outdated {
				log.Warning(fmt.Sprintf("%v of rootly is out! Please upgrade.", v))
			} else if err == nil {
				log.Success(false, "You are on the latest version of rootly.")
			}
			fmt.Println(version)
		} else {
			err := cmd.Help()
			if err != nil {
				log.Fatal(log.CtxErr{
					Context: "Failed to display help",
					Error:   err,
				})
			}
			os.Exit(0)
		}
	},
}

// Execute the root command
func Execute() {
	rootCmd.Flags().Bool(
		"version",
		false,
		"Get the current version of rootly and check for an update",
	)

	if err := rootCmd.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "unknown flag:") {
			os.Exit(1)
		}
		log.Fatal(log.CtxErr{
			Context: "Failed to execute root command",
			Error:   err,
		})
	}
}
