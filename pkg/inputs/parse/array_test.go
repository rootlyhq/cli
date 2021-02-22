package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	tt := []struct {
		str      string
		expected []string
	}{
		{
			str:      "Hello, World!",
			expected: []string{"Hello", "World!"},
		},
		{
			str:      "Hello,World!",
			expected: []string{"Hello", "World!"},
		},
		{
			str:      "Hello",
			expected: []string{"Hello"},
		},
		{
			str:      "",
			expected: []string{},
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, Array(test.str))
	}
}
