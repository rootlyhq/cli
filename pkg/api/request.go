package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/rootly.go"
)

type pulseData struct {
	Data struct {
		Type       string       `json:"type"`
		Attributes rootly.Pulse `json:"attributes"`
	} `json:"data"`
}

// Create a pulse on rootly.io
func CreatePulse(
	pulse rootly.Pulse,
	client *rootly.Client,
	secProvider *securityprovider.SecurityProviderBearerToken,
) log.CtxErr {
	log.Info("Creating pulse with summary of", pulse.Summary)

	// Marshaling data
	data, err := json.Marshal(pulseData{
		Data: struct {
			Type       string       "json:\"type\""
			Attributes rootly.Pulse "json:\"attributes\""
		}{
			Type:       "pulses",
			Attributes: pulse,
		},
	})
	if err != nil {
		return log.CtxErr{
			Context: "Failed to marshal data for creating a pulse",
			Error:   err,
		}
	}

	// Creating request
	req, err := rootly.NewCreatePulseRequestWithBody(
		serverName,
		"application/vnd.api+json",
		strings.NewReader(string(data)),
	)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to create pulse",
			Error:   err,
		}
	}

	// Intercept the request
	err = secProvider.Intercept(req.Context(), req)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to intercept request to inject auth header",
			Error:   err,
		}
	}

	// Make the request
	resp, err := client.Client.Do(req)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to make request to create pulse",
			Error:   err,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return log.NewErr("Failed to create pulse with exit code " + resp.Status)
	}

	log.Success("Created pulse with summary of", pulse.Summary)
	return log.CtxErr{}
}
