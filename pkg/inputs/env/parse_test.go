package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArray(t *testing.T) {
	tt := []struct {
		str      string
		expected []string
	}{
		{
			str:      "8,3,2,1",
			expected: []string{"8", "3", "2", "1"},
		},
		{
			str:      "",
			expected: []string{},
		},
		{
			str:      "\"8, 2\",3",
			expected: []string{"8, 2", "3"},
		},
	}

	for _, test := range tt {
		result, err := parseArray(test.str)

		assert.NoError(t, err.Error)
		assert.Equal(t, test.expected, result)
	}
}
