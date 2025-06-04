package inputs

import (
	"strings"

	"github.com/rootlyhq/cli/pkg/log"
)

// Convert an array to a simple map separated with |#|
func convertToSimpleMapArray(array []string) ([]map[string]string, log.CtxErr) {
	var finalVals []map[string]string

	for _, val := range array {
		twoVals := strings.Split(val, "=")

		if len(twoVals) == 1 {
			return []map[string]string{}, log.NewErr("Invalid format for '" + val + "'. Expected format: key=value")
		}
		if len(twoVals) >= 3 {
			return []map[string]string{}, log.NewErr("Invalid format for '" + val + "'. Too many '=' characters. Expected format: key=value")
		}
		if len(twoVals) == 2 && twoVals[1] == "" {
			return []map[string]string{}, log.NewErr("Invalid format for '" + val + "'. Value cannot be empty. Expected format: key=value")
		}

		finalVals = append(
			finalVals,
			map[string]string{
				"key":   strings.ReplaceAll(strings.ToLower(twoVals[0]), " ", "_"),
				"value": twoVals[1],
			},
		)
	}

	return finalVals, log.CtxErr{}
}
