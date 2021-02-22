package flags

import (
	"github.com/rootly-io/cli/pkg/inputs/parse"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/spf13/cobra"
)

// Error template for a failure to get the flag
func failedToGetErr(name string, err error) log.CtxErr {
	return log.CtxErr{
		Context: "Failed to get flag called " + name,
		Error:   err,
	}
}

// Get the value of a string flag
func GetString(name string, cmd *cobra.Command) (string, log.CtxErr) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return "", failedToGetErr(name, err)
	}
	return val, log.CtxErr{}
}

// Get the value of an array flag
func GetArray(name string, cmd *cobra.Command) ([]string, log.CtxErr) {
	str, err := cmd.Flags().GetString(name)
	if err != nil {
		return []string{}, failedToGetErr(name, err)
	}
	return parse.Array(str), log.CtxErr{}
}
