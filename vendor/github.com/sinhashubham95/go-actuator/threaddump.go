package actuator

import (
	"bytes"
	"errors"
	"net/http"
)

func getThreadDump() ([]byte, error) {
	var buffer bytes.Buffer
	profile := pprofLookupFunction(goRoutinesKey)
	if profile == nil {
		return nil, errors.New(profileNotFoundError)
	}
	err := profile.WriteTo(&buffer, 1)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// handleThreadDump is the handler to get the thread dump
func handleThreadDump(writer http.ResponseWriter, _ *http.Request) {
	body, err := getThreadDump()
	if err != nil {
		// some error occurred
		// send the error in the response
		sendStringResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
	// now once we have the correct response
	writer.Header().Add(contentTypeHeader, textStringContentType)
	_, _ = writer.Write(body)
}
