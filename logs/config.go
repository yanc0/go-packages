package logs

var (
	configActivateTimestamp bool
	configSeverity          string
	configFormat            string
	configDefaultFields     map[string]interface{}
)

func init() {
	ConfigActivateTimestamp(true)
	ConfigSeverity("info")
	ConfigFormat("text")
	ConfigDefaultFields(map[string]interface{}{})
}

// DefaultFields for each log generated
func DefaultFields(fields map[string]interface{}) {
	configDefaultFields = make(map[string]interface{})
	for key, value := range fields {
		configDefaultFields[key] = value
	}
}

// ConfigActivateTimestamp for log generation
func ConfigActivateTimestamp(timestamp bool) {
	configActivateTimestamp = timestamp
}

// ConfigSeverity for log generation
func ConfigSeverity(severity string) {
	configSeverity = severity
}

// ConfigFormat for log generation
func ConfigFormat(format string) {
	configFormat = format
}

// ConfigDefaultFields for log generation
func ConfigDefaultFields(fields map[string]interface{}) {
	configDefaultFields = fields
}
