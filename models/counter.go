package models

// CounterResponse is the response for the counter request
type CounterResponse struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}
