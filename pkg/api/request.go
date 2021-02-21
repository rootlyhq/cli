package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/cli/pkg/models"
	"github.com/rootly-io/rootly-go"
)

// Create a pulse on rootly.io
func CreatePulse(
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

	// Handeling a failed request
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return log.CtxErr{
				Context: "Failed to read response from failed request",
				Error:   err,
			}
		}

		if resp.StatusCode == http.StatusCreated {
			// Getting json data to see if serviceids or environmentids failed
			possibleErr := "The given environments or services don't exist in rootly."

			var jsonResp rootly.NewPulse
			err = json.Unmarshal(body, &jsonResp)
			if err != nil {
				return log.CtxErr{
					Context: "Failed to parse json from failed request response. " + possibleErr,
					Error:   err,
				}
			}

			// If there is an invalid service id
			if jsonResp.Data.Attributes.ServiceIds == nil && len(pulse.ServiceIds) != 0 {
				if len(pulse.ServiceIds) == 1 {
					return log.NewErr(pulse.ServiceIds[0] + " is not a valid service.")
				}
				return log.NewErr(
					"Please check your list of services. One is not valid: " + strings.Join(
						pulse.ServiceIds,
						", ",
					),
				)
			}

			// If there is an invalid environment id
			if jsonResp.Data.Attributes.EnvironmentIds == nil &&
				len(pulse.EnvironmentIds) != 0 {
				if len(pulse.EnvironmentIds) == 1 {
					return log.NewErr(pulse.EnvironmentIds[0] + " is not a valid environment.")
				}
				return log.NewErr(
					"Please check your list of environments. One is not valid: " + strings.Join(
						pulse.EnvironmentIds,
						", ",
					),
				)
			}
		}

		return log.NewErr(
			fmt.Sprintf(
				"Failed to create pulse with exit code %v\n\nPayload:\n%v\n\nResponse:\n%v",
				resp.Status,
				data,
				string(body),
			),
		)
	}

	log.Success("Created a pulse with the following values:", fmtPulse)
	return log.CtxErr{}
}
