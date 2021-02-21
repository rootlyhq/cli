package api

import (
	"encoding/json"
	"time"

	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/rootly-go"
)

// Convert a Pulse to a rootly.Pulse
func convertPulse(pulse Pulse) (string, log.CtxErr) {
	// Converting Pulse.Labels to rootly.Pulse.Lables
	var labels []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	for _, label := range pulse.Labels {
		labels = append(labels, struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{
			Key:   label["key"],
			Value: label["value"],
		})
	}

	// Putting rootly.Pulse inside apiData
	var finalData apiData
	finalData.Data.Type = "pulses"

	// We must conditionally add labels because the rootly library uses pointers
	if len(labels) == 0 {
		finalData.Data.Attributes = rootly.Pulse{
			Summary:   &pulse.Summary,
			CreatedAt: convertTime(pulse.CreatedAt),
			UpdatedAt: convertTime(pulse.UpdatedAt),
		}
	} else {
		finalData.Data.Attributes = rootly.Pulse{
			Summary:   &pulse.Summary,
			CreatedAt: convertTime(pulse.CreatedAt),
			UpdatedAt: convertTime(pulse.UpdatedAt),
			Labels:    &labels,
		}
	}

	// Marshaling the data
	jsonData, err := jsonMarshal(finalData)
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

// Convert time if given a value
func convertTime(t time.Time) string {
	var (
		epoch     = time.Time{}.Format(time.RFC3339)
		converted = t.Format(time.RFC3339)
	)

	if epoch == converted {
		return ""
	}
	return converted
}
