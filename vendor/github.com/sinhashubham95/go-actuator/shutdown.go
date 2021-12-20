package actuator

import (
	"net/http"
	"os"
)

// handleShutdown is the handler function for the shutdown endpoint
func handleShutdown(http.ResponseWriter, *http.Request) {
	// passing code 0 here to gracefully shut down the application
	os.Exit(0)
}
