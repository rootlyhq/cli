package api

import (
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/rootly.go"
)

// Create a pulse on rootly.io
func CreatePulse(
	pulse Pulse,
	client *rootly.Client,
	secProvider *securityprovider.SecurityProviderBearerToken,
) log.CtxErr {
	log.Info("Creating pulse with summary of", pulse.Summary)

	// Converting the data
	data, errCtx := convertPulse(pulse)
	if errCtx.Error != nil {
		return errCtx
	}

	// Creating request
	req, err := rootly.NewCreatePulseRequestWithBody(
		serverName,
		"application/vnd.api+json",
		strings.NewReader(data),
	)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to create pulse" + "\n\n" + data,
			Error:   err,
		}
	}

	// Intercept the request
	err = secProvider.Intercept(req.Context(), req)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to intercept request to inject auth header" + "\n\n" + data,
			Error:   err,
		}
	}

	// Make the request
	resp, err := client.Client.Do(req)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to make request to create pulse" + "\n\n" + data,
			Error:   err,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return log.NewErr("Failed to create pulse with exit code " + resp.Status + "\n\n" + data)
	}

	log.Success("Created pulse with summary of", pulse.Summary)
	return log.CtxErr{}
}
