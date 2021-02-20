package env

import (
	"encoding/csv"
	"strings"

	"github.com/rootly-io/cli/pkg/log"
)

// Parse an array from an environment variable
func parseArray(str string) ([]string, log.CtxErr) {
	strReader := strings.NewReader(str)
	values, err := csv.NewReader(strReader).Read()
	if err != nil {
		return []string{}, log.CtxErr{Context: "Failed to get "}
	}
	return values, log.CtxErr{}
}
