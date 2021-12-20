package actuator

import (
	"net/http"
	"runtime"
)

// Set of linked build time variables for providing relevant information for the application
var (
	BuildStamp       string
	GitCommitAuthor  string
	GitCommitID      string
	GitCommitTime    string
	GitPrimaryBranch string
	GitURL           string
	HostName         string
	Username         string
)

func getInfo(config *Config) map[string]interface{} {
	return map[string]interface{}{
		applicationKey: map[string]string{
			EnvKey:     config.Env,
			nameKey:    config.Name,
			versionKey: config.Version,
		},
		gitKey: map[string]string{
			buildStampKey:       BuildStamp,
			gitCommitAuthorKey:  GitCommitAuthor,
			gitCommitIDKey:      GitCommitID,
			gitCommitTimeKey:    GitCommitTime,
			gitPrimaryBranchKey: GitPrimaryBranch,
			gitURLKey:           GitURL,
			hostNameKey:         HostName,
			usernameKey:         Username,
		},
		runtimeKey: map[string]interface{}{
			archKey:           runtime.GOARCH,
			osKey:             runtime.GOOS,
			portKey:           config.Port,
			runtimeVersionKey: runtime.Version(),
		},
	}
}

// getInfoHandler is used to get the handler function for the info endpoint
func getInfoHandler(config *Config) http.HandlerFunc {
	info, _ := encodeJSON(getInfo(config))
	return func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Add(contentTypeHeader, applicationJSONContentType)
		_, _ = writer.Write(info)
	}
}
