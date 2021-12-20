package actuator

import (
	"bytes"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// isValidEndpoint is used to check if the endpoint is valid or not
func isValidEndpoint(endpoint int) bool {
	for _, e := range AllEndpoints {
		if endpoint == e {
			return true
		}
	}
	return false
}

// encodeJSON is used to encode any type of data to byte array
func encodeJSON(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	e := jsoniter.NewEncoder(&buffer)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// getLastStringAfterDelimiter is used to return the string after last delimiter existence
// a/b/c with / returns c
// a returns a
func getLastStringAfterDelimiter(s, delim string) string {
	splits := strings.Split(s, delim)
	length := len(splits)
	if length == 0 {
		return ""
	}
	return splits[length-1]
}

// sendStringResponse is used to write string response to the output
func sendStringResponse(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	if body != "" {
		w.Header().Add(contentTypeHeader, textStringContentType)
		_, _ = w.Write([]byte(body))
	}
}
