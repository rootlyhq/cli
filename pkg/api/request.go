package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	"github.com/rootlyhq/cli/pkg/log"
	"github.com/rootlyhq/cli/pkg/models"
	"github.com/rootlyhq/rootly-go"
)

// Create a pulse on rootly.com
func CreatePulse(
	host string,
	pulse models.Pulse,
	client *rootly.Client,
	secProvider *securityprovider.SecurityProviderBearerToken,
) log.CtxErr {
	fmtPulse := log.FormatPulse(pulse)
	log.Info("Creating pulse with the following values:", fmtPulse)

	// Converting the data
	data, errCtx := convertPulse(pulse)
	if errCtx.Error != nil {
		return errCtx
	}

	// Creating request
	req, err := rootly.NewCreatePulseRequestWithBody(
		host,
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
	req.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to read response from failed request",
			Error:   err,
		}
	}

	// Handeling a failed request
	if resp.StatusCode != http.StatusCreated {
		return log.NewErr(
			fmt.Sprintf(
				"Failed to create pulse with exit code %v\n\nPayload:\n%v\n\nResponse:\n%v",
				resp.Status,
				data,
				string(body),
			),
		)
	}

	// Getting json data to see if serviceids or environmentids failed
	possibleErr := "The given environments or services don't exist in rootly."
	var jsonResp models.CreatedPulseResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return log.CtxErr{
			Context: "Failed to parse json from failed request response. " + possibleErr,
			Error:   err,
		}
	}

	// If there is an invalid service id
	services := jsonResp.Data.Attributes.Services
	if len(services) != len(pulse.ServiceIds) {
		okServices := []string{}
		for _, service := range services {
			okServices = append(okServices, service.Name)
		}
		return log.NewErr(
			"The following services are invalid: " + strings.Join(
				arrayMissing(pulse.ServiceIds, okServices),
				", ",
			),
		)
	}

	// If there is an invalid environment id
	environments := jsonResp.Data.Attributes.Environments
	if len(environments) != len(pulse.EnvironmentIds) {
		okEnvironments := []string{}
		for _, environment := range environments {
			okEnvironments = append(okEnvironments, environment.Name)
		}
		return log.NewErr(
			"The following environments are invalid: " + strings.Join(
				arrayMissing(pulse.EnvironmentIds, okEnvironments),
				", ",
			),
		)
	}

	log.Success(true, "Created a pulse with the following values:", fmtPulse)
	return log.CtxErr{}
}

// Get a list of all expected items that aren't in a given array
func arrayMissing(expected []string, array []string) []string {
	missing := []string{}
	for _, expectedItem := range expected {
		found := false
		for _, item := range array {
			if item == expectedItem {
				found = true
			}
		}
		if !found {
			missing = append(missing, expectedItem)
		}
	}
	return missing
}
