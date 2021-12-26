package constants

const (
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
	RequestCountKey            = "requestCount"
	RequestCountDefaultValue   = 6000
	RequestCountUsage          = "Total number of requests"
	PathKey                    = "path"
	GapTimeKey                 = "gapTime"
	GapTimeValue               = 100
	GapTimeUsage               = "number of seconds to wait before writing to websockets"
	PathDefaultValue           = "/somepath"
	PathUsage                  = "Specific url path"
	LifeTimeKey                = "lifeTime"
	LifeTimeDefaultValue       = 5
	LifeTimeUsage              = "the duration for which each connection remains"
	StrategyKey                = "strategy"
	StrategyDefaultValue       = ""
	StrategyUsage              = "custom strategy which need to be injected"
	MessageTextKey             = "messageText"
	MessageTextDefaultValue    = "test"
	MessageTextUsage           = "the default text which need to be sent"
	WriteTimeKey               = "writeTime"
	WriteTimeDefaultValue      = 1
	WriteTimeUsage             = "The gap time between continuous writes"
)
