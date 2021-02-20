package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertName(t *testing.T) {
	tt := []struct {
		flag     string
		expected string
	}{
		{
			flag:     "help",
			expected: "ROOTLY_CLI_HELP",
		},
		{
			flag:     "HELP",
			expected: "ROOTLY_CLI_HELP",
		},
		{
			flag:     "old-help",
			expected: "ROOTLY_CLI_OLD_HELP",
		},
		{
			flag:     "OLD-HELP",
			expected: "ROOTLY_CLI_OLD_HELP",
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, convertName(test.flag))
	}
}
