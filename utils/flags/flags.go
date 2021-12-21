package flags

import (
	"github.com/angel-one/ws-load-test/constants"
	flag "github.com/spf13/pflag"
)

var (
	serverPort     = flag.Int(constants.PortKey, constants.PortDefaultValue, constants.PortUsage)
	host           = flag.String(constants.HostKey, constants.HostDefaultValue, constants.HostUsage)
	protocol       = flag.String(constants.ProtocolKey, constants.ProtocolDefaultValue, constants.ProtocolUsage)
	requestCount   = flag.Int(constants.RequestCountKey, constants.RequestCountDefaultValue, constants.RequestCountUsage)
	gapTime        = flag.Int(constants.GapTimeKey, constants.GapTimeValue, constants.GapTimeUsage)
	lifeTime       = flag.Int(constants.LifeTimeKey, constants.LifeTimeDefaultValue, constants.LifeTimeUsage)
	strategy       = flag.String(constants.StrategyKey, constants.StrategyDefaultValue, constants.StrategyUsage)
	messageText    = flag.String(constants.MessageTextKey, constants.MessageTextDefaultValue, constants.MessageTextUsage)
	writeTime      = flag.Int(constants.WriteTimeKey, constants.WriteTimeDefaultValue, constants.WriteTimeUsage)
	path           = flag.String(constants.PathKey, constants.PathDefaultValue, constants.PathUsage)
	baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
		constants.BaseConfigPathUsage)
)

func init() {
	flag.Parse()
}

func BaseConfigPath() string {
	return *baseConfigPath
}

func Host() string {
	return *host
}

func Strategy() string {
	return *strategy
}

func MessageText() string {
	return *messageText
}

func WriteTime() int {
	return *writeTime
}

func LifeTime() int {
	return *lifeTime
}

func Protocol() string {
	return *protocol
}

func Path() string {
	return *path
}

func GapTime() int {
	return *gapTime
}

func Request() int {
	return *requestCount
}

func ServerPort() int {
	return *serverPort
}
