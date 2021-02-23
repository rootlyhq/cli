package inputs

// A certain piece of config
type ConfigPiece string

const (
	// Name for the summary config related to pulses
	PulseSummaryName ConfigPiece = "summary"
	// Name for the label config related to pulses
	PulseLabelsName ConfigPiece = "labels"
	// Name for the service config related to pulses
	PulseServicesName ConfigPiece = "services"
	// Name for the environments config related to pulses
	PulseEnvironmentsName ConfigPiece = "environments"
	// Name for the API key config
	ApiKeyName ConfigPiece = "api-key"
	// Name for the API host config
	ApiHostName ConfigPiece = "api-host"
)
