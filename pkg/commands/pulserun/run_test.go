package pulserun

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunProgram(t *testing.T) {
	// Only run unit tests in linux or macOS
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		tt := []struct {
			program          string
			args             []string
			expectingErr     bool
			expectedExitCode int
		}{
			{
				program:          "echo",
				args:             []string{"hello", "world"},
				expectingErr:     false,
				expectedExitCode: 0,
			},
			{
				program:          "cat",
				args:             []string{"hello", "world"},
				expectingErr:     false,
				expectedExitCode: 1,
			},
			{
				program:          "program_that_doesnt_exist",
				args:             []string{"hello", "world"},
				expectingErr:     true,
				expectedExitCode: 0,
			},
		}

		for _, test := range tt {
			exitCode, err := RunProgram(test.program, test.args)

			if test.expectingErr {
				assert.Error(t, err.Error)
			} else {
				assert.NoError(t, err.Error)
			}

			assert.Equal(t, test.expectedExitCode, exitCode)
		}
	}
}
