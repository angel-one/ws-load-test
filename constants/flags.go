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
	HostUsage                  = "Domain name of test host. Ex 172.31.25.37:8080"
	ProtocolKey                = "protocol"
	ProtocolDefaultValue       = "wss"
	ProtocolUsage              = "Connection type. ws or wss."
	RequestCountKey            = "requestCount"
	RequestCountDefaultValue   = 6000
	RequestCountUsage          = "Total number of requests to be established through out load test."
	PathKey                    = "path"
	GapTimeKey                 = "gapTime"
	GapTimeValue               = 100
	GapTimeUsage               = "Number of milli seconds to wait in between establishing continuous ws connections."
	PathDefaultValue           = "/somepath"
	PathUsage                  = "Specific url path"
	LifeTimeKey                = "lifeTime"
	LifeTimeDefaultValue       = 5
	LifeTimeUsage              = "The duration for which each connection remains"
	StrategyKey                = "strategy"
	StrategyDefaultValue       = ""
	StrategyUsage              = "Custom strategy which need to be injected. ping_pong in case of ping pong strategy. Can also contain a custom strategy"
	MessageTextKey             = "messageText"
	MessageTextDefaultValue    = "test"
	MessageTextUsage           = "The default text which need to be sent in case a basic strategy is used. This need to be passed in case of custom pr ping-pong strategy"
	WriteTimeKey               = "writeTime"
	WriteTimeDefaultValue      = 1
	WriteTimeUsage             = "The gap time between continuous writes.This is in seconds."
)
