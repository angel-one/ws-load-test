package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

// GET is used to make a get request with the provided details
func GET(url string, headers map[string]string) (*http.Response, error) {
	return GETWithTimeout(url, headers, 0)
}

// GETWithTimeout is used to make a get request with the provided details
// 0 timeout means default timeout will be used
func GETWithTimeout(url string, headers map[string]string, timeout time.Duration) (*http.Response, error) {
	return GETWithTimeoutAndRetries(url, headers, timeout, 0, 0, 0)
}

// GETWithTimeoutAndRetries is used to make a get request with the provided details
// 0 timeout means default timeout will be used
func GETWithTimeoutAndRetries(url string, headers map[string]string, timeout time.Duration,
	retryCount int, retryWaitTime time.Duration, retryMaxWaitTime time.Duration) (*http.Response, error) {
	// create a request
	request, err := getRequest(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request %w", err)
	}

	// now time to execute with retry and backoff
	return doWithTimeoutAndRetries(request, timeout, retryCount, retryWaitTime, retryMaxWaitTime)
}
