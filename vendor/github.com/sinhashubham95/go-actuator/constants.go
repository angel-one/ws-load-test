package actuator

// Common Constants
const (
	applicationKey      = "app"
	archKey             = "arch"
	buildStampKey       = "buildStamp"
	EnvKey              = "env"
	equals              = "="
	gitKey              = "git"
	gitCommitAuthorKey  = "commitAuthor"
	gitCommitIDKey      = "commitId"
	gitCommitTimeKey    = "commitTime"
	gitPrimaryBranchKey = "branch"
	gitURLKey           = "url"
	goRoutinesKey       = "goroutine"
	hostNameKey         = "hostName"
	nameKey             = "name"
	portKey             = "port"
	osKey               = "os"
	runtimeKey          = "runtime"
	runtimeVersionKey   = "runtimeVersion"
	usernameKey         = "username"
	versionKey          = "version"
	slash               = "/"
)

// Endpoints
const (
	envEndpoint        = "/env"
	infoEndpoint       = "/info"
	metricsEndpoint    = "/metrics"
	pingEndpoint       = "/ping"
	shutdownEndpoint   = "/shutdown"
	threadDumpEndpoint = "/threadDump"
)

// Response constants
const (
	contentTypeHeader          = "Content-Type"
	applicationJSONContentType = "application/json"
	textStringContentType      = "text/string"
)

// Error messages
const (
	methodNotAllowedError = "requested method is not allowed on the called endpoint"
	notFoundError         = "not found"
	profileNotFoundError  = "profile not found"
)
