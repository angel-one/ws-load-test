package actuator

import (
	"net/http"
	"os"
	"strings"
)

func getEnvironmentVariables(config *Config) map[string]string {
	variables := make(map[string]string)
	for _, e := range os.Environ() {
		keyValue := strings.SplitN(e, equals, 2)
		if len(keyValue) == 2 {
			variables[keyValue[0]] = keyValue[1]
		}
	}
	variables[EnvKey] = config.Env
	return variables
}

// getEnvHandler is used to provide the handler function for the env endpoint
func getEnvHandler(config *Config) http.HandlerFunc {
	variables, _ := encodeJSON(getEnvironmentVariables(config))
	return func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Add(contentTypeHeader, applicationJSONContentType)
		_, _ = writer.Write(variables)
	}
}
