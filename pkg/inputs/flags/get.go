package flags

import (
	"github.com/rootly-io/cli/pkg/inputs/parse"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/cli/pkg/models"
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
func GetString(name models.ConfigPiece, cmd *cobra.Command) (string, log.CtxErr) {
	strName := string(name)
	val, err := cmd.Flags().GetString(strName)
	if err != nil {
		return "", failedToGetErr(strName, err)
	}
	return val, log.CtxErr{}
}

// Get the value of a boolean flag
func GetBool(name models.ConfigPiece, cmd *cobra.Command) (bool, log.CtxErr) {
	strName := string(name)
	val, err := cmd.Flags().GetBool(strName)
	if err != nil {
		return false, failedToGetErr(strName, err)
	}
	return val, log.CtxErr{}
}

// Get the value of an array flag
func GetArray(name models.ConfigPiece, cmd *cobra.Command) ([]string, log.CtxErr) {
	strName := string(name)
	str, err := cmd.Flags().GetString(strName)
	if err != nil {
		return []string{}, failedToGetErr(strName, err)
	}
	return parse.Array(str), log.CtxErr{}
}
