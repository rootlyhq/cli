package api

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rootly-io/cli/pkg/log"
	"github.com/rootly-io/rootly-go"
)

// Generate a rootly client
func GenClient(host string) (*rootly.Client, log.CtxErr) {
	client, err := rootly.NewClient(host)
	if err != nil {
		return &rootly.Client{}, log.CtxErr{
			Context: "Failed to create rootly client",
			Error:   err,
		}
	}
	return client, log.CtxErr{}
}

// Get a security provider to wrap requests with
func GenSecurityProvider(
	apiKey string,
) (*securityprovider.SecurityProviderBearerToken, log.CtxErr) {
	provider, err := securityprovider.NewSecurityProviderBearerToken(apiKey)
	if err != nil {
		return &securityprovider.SecurityProviderBearerToken{}, log.CtxErr{
			Context: "Failed to create security provider with api key",
			Error:   err,
		}
	}
	return provider, log.CtxErr{}
}
