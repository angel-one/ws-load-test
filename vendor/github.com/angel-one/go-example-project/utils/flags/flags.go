package flags

import (
	"github.com/angel-one/go-example-project/constants"
	flag "github.com/spf13/pflag"
)

var (
	env            = flag.String(constants.EnvKey, constants.EnvDefaultValue, constants.EnvUsage)
	port           = flag.Int(constants.PortKey, constants.PortDefaultValue, constants.PortUsage)
	baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
		constants.BaseConfigPathUsage)
)

func init() {
	flag.Parse()
}

// Env is the application.yml runtime environment
func Env() string {
	return *env
}

// Port is the application.yml port number where the process will be started
func Port() int {
	return *port
}

// BaseConfigPath is the path that holds the configuration files
func BaseConfigPath() string {
	return *baseConfigPath
}
