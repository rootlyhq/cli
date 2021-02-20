package main

import (
	"github.com/Matt-Gleich/lumber"
	"github.com/rootly-io/cli/pkg/commands"
)

func main() {

	lumber.ErrNilCheck = false
	commands.Execute()

	// req, err := rootly.NewCreatePulseRequestWithBody("https://api.rootly.io", "application/vnd.api+json", strings.NewReader("{\"data\": {\"type\": \"pulses\",\"attributes\": {\"summary\": \"hello world\"}}}"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// err = bearerToken.Intercept(req.Context(), req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// _, err = client.Client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
