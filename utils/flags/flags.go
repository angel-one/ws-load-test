package flags

import (
	"github.com/angel-one/ws-load-test/constants"
	flag "github.com/spf13/pflag"
)

var (
	host           = flag.String(constants.HostKey, constants.HostDefaultValue, constants.HostUsage)
	protocol       = flag.String(constants.ProtocolKey, constants.ProtocolDefaultValue, constants.ProtocolUsage)
	request        = flag.Int(constants.RequestKey, constants.RequestDefaultValue, constants.RequestUsage)
	writeTime      = flag.Int(constants.WriteTimeKey, constants.WriteTimeValue, constants.WriteTimeUsage)
	holdTime       = flag.Int(constants.HoldTimeKey, constants.HoldTimeDefaultValue, constants.HoldTimeUsage)
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

func Protocol() string {
	return *protocol
}

func Path() string {
	return *path
}

func WriteTime() int {
	return *writeTime
}

func HoldTime() int {
	return *holdTime
}

func Request() int {
	return *request
}
