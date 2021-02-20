package api

import "time"

// Wrapper for some API data
type apiData struct {
	Data struct {
		Type       string      `json:"type"`
		Attributes interface{} `json:"attributes"`
	} `json:"data"`
}

// Structure for a pulse
type Pulse struct {
	Summary   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Labels    []map[string]string
}
