package httpclient

import (
	"github.com/angel-one/ws-load-test/constants"
	"github.com/angel-one/go-utils/log"
	"golang.org/x/net/publicsuffix"
	"net"
	"net/http"
	"net/http/cookiejar"
	"runtime"
	"time"
)

// Config is the set of configurable parameters for the http client
type Config struct {
	// ConnectTimeout is the maximum amount of time a dial will wait for
	// connect to complete. If Deadline is also set, it may fail
	// earlier.
	ConnectTimeout time.Duration `json:"connectTimeout"`
	// KeepAliveDuration specifies the interval between keep-alive
	// probes for an active network connection.
	KeepAliveDuration time.Duration `json:"KeepAliveDuration"`
	// MaxIdleConnections controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConnections int `json:"maxIdleConnections"`
	// IdleConnectionTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	IdleConnectionTimeout time.Duration `json:"idleConnectionTimeout"`
	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake.
	TLSHandshakeTimeout time.Duration `json:"tlsHandshakeTimeout"`
	// ExpectContinueTimeout specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers.
	ExpectContinueTimeout time.Duration `json:"expectContinueTimeout"`
	// Timeout specifies a time limit for requests made by this
	// Client.
	Timeout time.Duration `json:"timeout"`
}

var client *http.Client

func Init(config Config) error {
	log.Info(nil).Interface(constants.HTTPConfigKey, config).Msg("initializing http client")
	cookieJar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	dialer := &net.Dialer{
		Timeout:   config.ConnectTimeout,
		KeepAlive: config.KeepAliveDuration,
	}

	client = &http.Client{
		Jar: cookieJar,
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          config.MaxIdleConnections,
			IdleConnTimeout:       config.IdleConnectionTimeout,
			TLSHandshakeTimeout:   config.TLSHandshakeTimeout,
			ExpectContinueTimeout: config.ExpectContinueTimeout,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
		// global timeout value for all requests
		Timeout: config.Timeout,
	}
	return nil
}
