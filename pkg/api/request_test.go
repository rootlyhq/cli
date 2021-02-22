package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayMissing(t *testing.T) {
	tt := []struct {
		expected        []string
		array           []string
		expectedMissing []string
	}{
		{
			expected:        []string{"Hello", "World"},
			array:           []string{"Hello", "World"},
			expectedMissing: []string{},
		},
		{
			expected:        []string{"Hello", "World"},
			array:           []string{"Hello"},
			expectedMissing: []string{"World"},
		},
		{
			expected:        []string{},
			array:           []string{},
			expectedMissing: []string{},
		},
		{
			expected:        []string{"Hello World"},
			array:           []string{"Hello World"},
			expectedMissing: []string{},
		},
		{
			expected:        []string{"h"},
			array:           []string{},
			expectedMissing: []string{"h"},
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expectedMissing, arrayMissing(test.expected, test.array))
	}
}
