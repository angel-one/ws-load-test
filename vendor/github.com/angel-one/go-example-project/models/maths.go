package models

import "github.com/angel-one/go-example-project/jobs"

// FullNameRequest is the request body for the full name api
type MathsRequest struct {
	A  int    `json:"a"`
	B  int    `json:"b"`
	Op string `json:"op"`
}

type MathsResult struct {
	State   jobs.JobState `json:"state_code"` // 1 is Active 6 is Completed
	Message string        `json:"status"`
	Result  int           `json:"res"`
}

type MathsJobResponse struct {
	Success bool   `json:"success"`
	JobId   string `json:"res"`
}
