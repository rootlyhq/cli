package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rootly-io/rootly.go"
)

func main() {

	bearerToken, err := securityprovider.NewSecurityProviderBearerToken("AUTH TOKEN")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := rootly.NewClient("https://api.rootly.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req, err := rootly.NewCreatePulseRequestWithBody("https://api.rootly.io", "application/vnd.api+json", strings.NewReader("{\"data\": {\"type\": \"pulses\",\"attributes\": {\"summary\": \"hello world\"}}}"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = bearerToken.Intercept(req.Context(), req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = client.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
