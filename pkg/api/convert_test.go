package api

import (
	"testing"
	"time"

	"github.com/rootly-io/cli/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestConvertPulse(t *testing.T) {

	tt := []struct {
		pulse    models.Pulse
		expected string
	}{
		{
			pulse:    models.Pulse{Summary: "Hello World!"},
			expected: "{\"data\":{\"attributes\":{\"ended_at\":null,\"environment_ids\":null,\"service_ids\":null,\"started_at\":null,\"summary\":\"Hello World!\"},\"type\":\"pulses\"}}",
		},
		{
			pulse: models.Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "platform",
						"value": "osx",
					},
				},
			},
			expected: "{\"data\":{\"attributes\":{\"ended_at\":null,\"environment_ids\":null,\"labels\":[{\"key\":\"platform\",\"value\":\"osx\"}],\"service_ids\":null,\"started_at\":null,\"summary\":\"Hello World!\"},\"type\":\"pulses\"}}",
		},
		{
			pulse: models.Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "platform",
						"value": "osx",
					},
				},
				StartedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "{\"data\":{\"attributes\":{\"ended_at\":null,\"environment_ids\":null,\"labels\":[{\"key\":\"platform\",\"value\":\"osx\"}],\"service_ids\":null,\"started_at\":\"2001-01-01T01:01:01.000000001Z\",\"summary\":\"Hello World!\"},\"type\":\"pulses\"}}",
		},
		{
			pulse: models.Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "platform",
						"value": "osx",
					},
				},
				EndedAt:   time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
				StartedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "{\"data\":{\"attributes\":{\"ended_at\":\"2001-01-01T01:01:01.000000001Z\",\"environment_ids\":null,\"labels\":[{\"key\":\"platform\",\"value\":\"osx\"}],\"service_ids\":null,\"started_at\":\"2001-01-01T01:01:01.000000001Z\",\"summary\":\"Hello World!\"},\"type\":\"pulses\"}}",
		},
		{
			pulse: models.Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "platform",
						"value": "osx",
					},
					{
						"key":   "Exit Code",
						"value": "1",
					},
				},
				EndedAt:        time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
				StartedAt:      time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
				ServiceIds:     []string{"elasticsearch-prod"},
				EnvironmentIds: []string{"production", "staging"},
			},
			expected: "{\"data\":{\"attributes\":{\"ended_at\":\"2001-01-01T01:01:01.000000001Z\",\"environment_ids\":[\"production\",\"staging\"],\"labels\":[{\"key\":\"platform\",\"value\":\"osx\"},{\"key\":\"Exit Code\",\"value\":\"1\"}],\"service_ids\":[\"elasticsearch-prod\"],\"started_at\":\"2001-01-01T01:01:01.000000001Z\",\"summary\":\"Hello World!\"},\"type\":\"pulses\"}}",
		},
	}

	for _, test := range tt {
		converted, err := convertPulse(test.pulse)
		assert.NoError(t, err.Error)

		assert.Equal(t, test.expected, converted)
	}
}
