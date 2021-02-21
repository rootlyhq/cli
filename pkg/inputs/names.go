package inputs

// A certain piece of config
type ConfigPiece string

const (
	// Name for the summary config related to pulses
	PulseSummaryName ConfigPiece = "summary"
	// Name for the label config related to pulses
	PulseLabelsName ConfigPiece = "label"
	// Name for the service config related to pulses
	PulseServicesName ConfigPiece = "service"
	// Name for the environments config related to pulses
	PulseEnvironmentsName ConfigPiece = "environment"
	// Name for the API key config
	ApiKeyName ConfigPiece = "api-key"
)
