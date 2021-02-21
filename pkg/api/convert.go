package api

import (
	"encoding/json"

	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/rootly-go"
)

// Convert a Pulse to a json version of rootly.NewPulse
func convertPulse(pulse Pulse) (string, log.CtxErr) {
	// Converting Pulse.Labels to rootly.NewPulse.Lables
	labels := []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{}
	for _, label := range pulse.Labels {
		labels = append(labels, struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{
			Key:   label["key"],
			Value: label["value"],
		})
	}

	// Putting data into rootly.NewPulse
	// We need to add optional data conditionally because the
	// rootly library uses pointers for everything
	var data rootly.NewPulse
	data.Data.Type = "pulses"
	data.Data.Attributes.Summary = &pulse.Summary
	if !pulse.EndedAt.IsZero() {
		data.Data.Attributes.EndedAt = &pulse.EndedAt
	}
	if !pulse.StartedAt.IsZero() {
		data.Data.Attributes.StartedAt = &pulse.StartedAt
	}
	if len(pulse.ServiceIds) != 0 {
		data.Data.Attributes.ServiceIds = &pulse.ServiceIds
	}
	if len(pulse.EnvironmentIds) != 0 {
		data.Data.Attributes.EnvironmentIds = &pulse.EnvironmentIds
	}
	if len(labels) != 0 {
		data.Data.Attributes.Labels = &labels
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
