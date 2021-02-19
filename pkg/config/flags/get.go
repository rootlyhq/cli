package flags

import (
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

// Error template for a failure to get the flag
func failedToGetErr(flag string, err error) log.CtxErr {
	return log.CtxErr{
		Context: "Failed to get flag called " + flag,
		Error:   err,
	}
}

// Error template for a blank value
func failedToProvideValue(flag string) log.CtxErr {
	return log.NewErr("Please provide a value for the " + flag + " flag")
}

// Get the value of a required string flag
func GetRequiredString(flag string, cmd *cobra.Command) (string, log.CtxErr) {
	val, err := cmd.Flags().GetString(flag)
	if err != nil {
		return "", failedToGetErr(flag, err)
	}

	if val == "" {
		return "", failedToProvideValue(flag)
	}

	return val, log.CtxErr{}
}

// Get the value of a required string  array flag
func GetRequiredStringArray(flag string, cmd *cobra.Command) ([]string, log.CtxErr) {
	val, err := cmd.Flags().GetStringArray(flag)
	if err != nil {
		return []string{}, failedToGetErr(flag, err)
	}

	if len(val) == 0 {
		return []string{}, failedToProvideValue(flag)
	}

	return val, log.CtxErr{}
}
