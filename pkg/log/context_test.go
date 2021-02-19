package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErr(t *testing.T) {
	tt := []struct {
		context  string
		expected CtxErr
	}{
		{
			context: "Testing Testing",
			expected: CtxErr{
				Context: "Testing Testing",
				Error:   errors.New("Testing Testing"),
			},
		},
		{
			context: "Hello World",
			expected: CtxErr{
				Context: "Hello World",
				Error:   errors.New("Hello World"),
			},
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, NewErr(test.context))
	}
}
