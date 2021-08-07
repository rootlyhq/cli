package inputs

import (
	"strings"

	"github.com/rootlyhq/cli/pkg/log"
)

// Convert an array to a simple map separated with |#|
func convertToSimpleMapArray(array []string) []map[string]string {
	var finalVals []map[string]string

	for _, val := range array {
		twoVals := strings.Split(val, "=")

		if len(twoVals) >= 3 {
			log.Warning(twoVals[0], "has more than one correspoding value. It will be ignored")
			continue
		}
		if len(twoVals) == 2 && twoVals[1] == "" {
			log.Warning(twoVals[0], "has no corresponding value. It will be ignored")
			continue
		}

		finalVals = append(
			finalVals,
			map[string]string{
				"key":   strings.ReplaceAll(strings.ToLower(twoVals[0]), " ", "_"),
				"value": twoVals[1],
			},
		)
	}

	return finalVals
}
