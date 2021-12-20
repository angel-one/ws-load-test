package actuator

import "runtime/pprof"

// variables for mocking in case of unit testing
var (
	encodeJSONFunction  = encodeJSON
	pprofLookupFunction = pprof.Lookup
)
