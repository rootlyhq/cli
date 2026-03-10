package api

import (
	"encoding/json"
	"time"

	"github.com/oapi-codegen/nullable"
	"github.com/rootlyhq/cli/pkg/log"
	"github.com/rootlyhq/cli/pkg/models"
	"github.com/rootlyhq/rootly-go"
)

// Convert a map to nullable API objects
func convertObject(maps []map[string]string) []nullable.Nullable[struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}] {
	var objects []nullable.Nullable[struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}]
	for _, mapData := range maps {
		objects = append(objects, nullable.NewNullableWithValue(struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{
			Key:   mapData["key"],
			Value: mapData["value"],
		}))
	}
	return objects
}

// Convert a Pulse to a json version of rootly.NewPulse
func convertPulse(pulse models.Pulse) (string, log.CtxErr) {
	// Putting data into rootly.NewPulse
	// We need to add optional data conditionally because the
	// rootly library uses nullable types for optional fields
	var data rootly.NewPulse
	data.Data.Type = "pulses"
	data.Data.Attributes.Summary = pulse.Summary
	data.Data.Attributes.Source = nullable.NewNullableWithValue(pulse.Source)
	if !pulse.EndedAt.IsZero() {
		data.Data.Attributes.EndedAt = nullable.NewNullableWithValue[time.Time](pulse.EndedAt)
	}
	if !pulse.StartedAt.IsZero() {
		data.Data.Attributes.StartedAt = nullable.NewNullableWithValue[time.Time](pulse.StartedAt)
	}
	if len(pulse.ServiceIds) != 0 {
		data.Data.Attributes.ServiceIDs = nullable.NewNullableWithValue(pulse.ServiceIds)
	}
	if len(pulse.EnvironmentIds) != 0 {
		data.Data.Attributes.EnvironmentIDs = nullable.NewNullableWithValue(pulse.EnvironmentIds)
	}
	labels := convertObject(pulse.Labels)
	if len(labels) != 0 {
		data.Data.Attributes.Labels = labels
	}
	refs := convertObject(pulse.Refs)
	if len(refs) != 0 {
		data.Data.Attributes.Refs = refs
	}

	// Marshaling the data
	jsonData, err := jsonMarshal(data)
	if err.Error != nil {
		return "", err
	}

	return jsonData, log.CtxErr{}
}

// Light wrapper around json.Marhal for consistent errors
func jsonMarshal(data interface{}) (string, log.CtxErr) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", log.CtxErr{
			Context: "Failed to marshal data for creating a pulse",
			Error:   err,
		}
	}
	return string(str), log.CtxErr{}
}
