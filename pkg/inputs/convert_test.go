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

	tt := []struct {
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
				fmt.Sprintf("Exit Code%v", delimiter),
			},
			expected: []map[string]string{{keyName: "exit_code", valName: "2"}},
		},
		{
			array: []string{
				fmt.Sprintf("Exit Code%v2", delimiter),
				fmt.Sprintf("Name%v", delimiter),
			},
			expected: []map[string]string{{keyName: "exit_code", valName: "2"}},
		},
		{
			array: []string{
				fmt.Sprintf("Exit Code%v2", delimiter),
				fmt.Sprintf("Name%vHarry%vPotter", delimiter, delimiter),
			},
			expected: []map[string]string{{keyName: "exit_code", valName: "2"}},
		},
		{
			array: []string{
				fmt.Sprintf("Exit Code%v2", delimiter),
				fmt.Sprintf("Platform%vOSX", delimiter),
			},
			expected: []map[string]string{
				{keyName: "exit_code", valName: "2"},
				{keyName: "platform", valName: "osx"},
			},
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, convertToSimpleMapArray(test.array))
	}
}
