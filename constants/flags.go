package constants

// Flag constants
const (
	EnvKey                     = "env"
	EnvDefaultValue            = ""
	EnvUsage                   = "application.yml runtime environment"
	PortKey                    = "port"
	PortDefaultValue           = 8080
	PortUsage                  = "application.yml port"
	BaseConfigPathKey          = "base-config-path"
	BaseConfigPathDefaultValue = "resources"
	BaseConfigPathUsage        = "path to folder that stores your configurations"
	HostKey                    = "host"
	HostDefaultValue           = "example.com"
	HostUsage                  = "Domain name of test host"
	ProtocolKey                = "protocol"
	ProtocolDefaultValue       = "wss"
	ProtocolUsage              = "Connection type"
	RequestKey                 = "request"
	RequestDefaultValue        = 6000
	RequestUsage               = "Total number of requests"
	WriteTimeKey               = "wtime"
	WriteTimeValue             = 100
	WriteTimeUsage             = "number of seconds to wait before writing to websockets"
	HoldTimeKey                = "htime"
	HoldTimeDefaultValue       = 30
	HoldTimeUsage              = "number of milliseconds to wait before creating new websocket connection"
	PathDefaultValue           = "/somepath"
	PathUsage                  = "Specific url path"
)
