package constants

// config names
const (
	LoggerConfig      = "logger"
	MoxyConfig        = "moxy"
	ApplicationConfig = "application"
	DatabaseConfig    = "database"
	JobsConfig        = "jobs"
	CounterConfig     = "counter"
)

// config keys
const (
	LogLevelConfigKey                         = "level"
	URLConfigKey                              = "url"
	HTTPConnectTimeoutInMillisKey             = "http.connectTimeoutInMillis"
	HTTPKeepAliveDurationInMillisKey          = "http.keepAliveDurationInMillis"
	HTTPMaxIdleConnectionsKey                 = "http.maxIdleConnections"
	HTTPIdleConnectionTimeoutInMillisKey      = "http.idleConnectionTimeoutInMillis"
	HTTPTlsHandshakeTimeoutInMillisKey        = "http.tlsHandshakeTimeoutInMillis"
	HTTPExpectContinueTimeoutInMillisKey      = "http.expectContinueTimeoutInMillis"
	HTTPTimeoutInMillisKey                    = "http.timeoutInMillis"
	DatabaseServerConfigKey                   = "server"
	DatabasePortConfigKey                     = "port"
	DatabaseNameConfigKey                     = "name"
	DatabaseUsernameConfigKey                 = "username"
	DatabasePasswordConfigKey                 = "password"
	DatabaseMaxOpenConnectionsKey             = "maxOpenConnections"
	DatabaseMaxIdleConnectionsKey             = "maxIdleConnections"
	DatabaseConnectionMaxLifetimeInSecondsKey = "connectionMaxLifetimeInSeconds"
	DatabaseConnectionMaxIdleTimeInSecondsKey = "connectionMaxIdleTimeInSeconds"
	CounterQueryTimeoutInMillisKey            = "queryTimeoutInMillis"
)

// Jobs Config
const (
	RedisConnectionString = "redisUrl"
	JobsTTLInHrs = "jobsRetentionTimeInHours"
	NumberOfWorkers = "numberOfWorkers"
)
