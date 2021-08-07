package log

import (
	"testing"
	"time"

	"github.com/rootlyhq/cli/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestFormatPulse(t *testing.T) {
	tt := []struct {
		pulse    models.Pulse
		expected string
	}{
		{
			pulse: models.Pulse{
				Summary: "Basic Pulse",
			},
			expected: "\n\n  Summary:      Basic Pulse\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				StartedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 01:01 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				EndedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 01:01 UTC\n  Services:     None\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				ServiceIds: []string{"Service 1"},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     Service 1\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				ServiceIds: []string{"Service 1", "Service 2"},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     Service 1, Service 2\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				EnvironmentIds: []string{"Environment 1", "Service 2"},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: Environment 1, Service 2\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				EnvironmentIds: []string{"Environment 1", "Environment 2"},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: Environment 1, Environment 2\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				Labels: []map[string]string{
					{"key": "Exit Code", "value": "0"},
				},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: None\n  Labels:       \n    Exit Code = 0\n\n  Source:       None\n  Refs:         None",
		},
		{
			pulse: models.Pulse{
				Labels: []map[string]string{
					{"key": "Exit Code", "value": "0"},
					{"key": "Version", "value": "1.0.0"},
				},
				Refs: []map[string]string{
					{"key": "sha", "value": "cd62148cbc5eb42168fe99fdb50a364e12b206ac"},
					{"key": "image", "value": "registry.rootly.io/rootly/my-service:cd6214"},
				},
			},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: None\n  Labels:       \n    Exit Code = 0\n    Version = 1.0.0\n\n  Source:       None\n  Refs:         \n    sha = cd62148cbc5eb42168fe99fdb50a364e12b206ac\n    image = registry.rootly.io/rootly/my-service:cd6214",
		},
		{
			pulse:    models.Pulse{},
			expected: "\n\n  Summary:      None\n  Started At:   01 Jan 01 00:00 UTC\n  Ended At      01 Jan 01 00:00 UTC\n  Services:     None\n  Environments: None\n  Labels:       None\n  Source:       None\n  Refs:         None",
		},
	}

	for _, test := range tt {
		assert.Equal(t, test.expected, FormatPulse(test.pulse))
	}
}
