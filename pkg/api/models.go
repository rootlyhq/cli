package api

import "time"

// Structure for a pulse
type Pulse struct {
	Summary        string
	StartedAt      time.Time
	EndedAt        time.Time
	ServiceIds     []string
	EnvironmentIds []string
	Labels         []map[string]string
}
