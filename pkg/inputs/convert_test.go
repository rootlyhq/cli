package inputs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToSimpleMapArray(t *testing.T) {
	const (
		keyName   = "key"
		valName   = "value"
		delimiter = "="
	)

	// Test valid cases
	validTests := []struct {
		array    []string
		expected []map[string]string
	}{
		{
			array:    []string{fmt.Sprintf("Exit Code%v2", delimiter)},
			expected: []map[string]string{{keyName: "exit_code", valName: "2"}},
		},
		{
			array: []string{
				fmt.Sprintf("Exit Code%v2", delimiter),
				fmt.Sprintf("Platform%vOSX", delimiter),
			},
			expected: []map[string]string{
				{keyName: "exit_code", valName: "2"},
				{keyName: "platform", valName: "OSX"},
			},
		},
	}

	for _, test := range validTests {
		result, err := convertToSimpleMapArray(test.array)
		assert.Nil(t, err.Error)
		assert.Equal(t, test.expected, result)
	}

	// Test invalid cases - should return errors
	invalidTests := []struct {
		array       []string
		description string
	}{
		{
			array:       []string{"Exit Code=2", "Exit Code="},
			description: "empty value",
		},
		{
			array:       []string{"Exit Code=2", "Name=Harry=Potter"},
			description: "multiple equals signs",
		},
		{
			array:       []string{"Exit Code=2", "InvalidFormat"},
			description: "missing equals sign",
		},
	}

	for _, test := range invalidTests {
		result, err := convertToSimpleMapArray(test.array)
		assert.NotNil(t, err.Error, "Expected error for "+test.description)
		assert.Empty(t, result, "Expected empty result for "+test.description)
	}
}
