package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertPulse(t *testing.T) {

	tt := []struct {
		pulse    Pulse
		expected string
	}{
		{
			pulse:    Pulse{Summary: "Hello World!"},
			expected: "{\"data\":{\"type\":\"pulses\",\"attributes\":{\"created_at\":\"\",\"labels\":null,\"summary\":\"Hello World!\",\"updated_at\":\"\"}}}",
		},
		{
			pulse: Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "Name",
						"value": "Harry Potter",
					},
				},
			},
			expected: "{\"data\":{\"type\":\"pulses\",\"attributes\":{\"created_at\":\"\",\"labels\":[{\"key\":\"Name\",\"value\":\"Harry Potter\"}],\"summary\":\"Hello World!\",\"updated_at\":\"\"}}}",
		},
		{
			pulse: Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "Name",
						"value": "Harry Potter",
					},
				},
				UpdatedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "{\"data\":{\"type\":\"pulses\",\"attributes\":{\"created_at\":\"\",\"labels\":[{\"key\":\"Name\",\"value\":\"Harry Potter\"}],\"summary\":\"Hello World!\",\"updated_at\":\"2001-01-01T01:01:01Z\"}}}",
		},
		{
			pulse: Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "Name",
						"value": "Harry Potter",
					},
				},
				CreatedAt: time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			expected: "{\"data\":{\"type\":\"pulses\",\"attributes\":{\"created_at\":\"2001-01-01T01:01:01Z\",\"labels\":[{\"key\":\"Name\",\"value\":\"Harry Potter\"}],\"summary\":\"Hello World!\",\"updated_at\":\"\"}}}",
		},
		{
			pulse: Pulse{
				Summary: "Hello World!",
				Labels: []map[string]string{
					{
						"key":   "Name",
						"value": "Harry Potter",
					},
					{
						"key":   "Exit Code",
						"value": "1",
					},
				},
			},
			expected: "{\"data\":{\"type\":\"pulses\",\"attributes\":{\"created_at\":\"\",\"labels\":[{\"key\":\"Name\",\"value\":\"Harry Potter\"},{\"key\":\"Exit Code\",\"value\":\"1\"}],\"summary\":\"Hello World!\",\"updated_at\":\"\"}}}",
		},
	}

	for _, test := range tt {
		converted, err := convertPulse(test.pulse)
		assert.NoError(t, err.Error)

		assert.Equal(t, test.expected, converted)
	}
}
