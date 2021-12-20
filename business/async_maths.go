package business

import (
	"encoding/json"
	"errors"
	"github.com/angel-one/go-example-project/models"
	"log"

	"github.com/angel-one/go-example-project/jobs"
)

const (
	AsyncMultiply = "math:multiply"
	AsyncAdd      = "math:add"
	MathsTaskType = "math"
)

type MathPayload struct {
	A int
	B int
}

var jobsController *jobs.JobController

func InitMaths(j *jobs.JobController) {
	jobsController = j
	jobsController.AddHandler(MathsTaskType, MathsHandler)
}

func SubmitMathsJob(request *models.MathsRequest) *models.MathsJobResponse {
	// TODO : some validation & transations of request

	// Create a async response
	var response models.MathsJobResponse

	// Construst args
	args := make(map[string]interface{})
	args["A"] = request.A
	args["B"] = request.B
	args["op"] = request.Op

	jobId, err := jobsController.CreateJob(MathsTaskType, args)
	if err != nil {
		log.Printf("error while submitting job : %v\n", err)
		response.Success = false

	} else {
		response.Success = true
		response.JobId = jobId
	}

	return &response
}
func GetMathJobStatus(id string) (*models.MathsResult, error) {
	// TODO ideally this should do some transalation of JobStatus to business model
	taskInfo, err := jobsController.GetStatus(id)
	if err != nil {
		return nil, err
	}

	// the value to return
	var result models.MathsResult

	result.State = taskInfo.State // todo needs transalation
	result.Message = taskInfo.Status
	var jobRes jobs.JobResult
	json.Unmarshal(taskInfo.Result, &jobRes)
	res, ok := jobRes.Result.(float64) // JSON treats all nos as float64
	if !ok {
		result.State = jobs.Failed
		return nil, errors.New("unable to deser math result")
	}
	result.Result = int(res)

	return &result, nil

}

func MathsHandler(args map[string]interface{}) *jobs.JobResult {
	var jobRes jobs.JobResult

	aRaw, ok := args["A"]
	if !ok {
		jobRes.Success = false
		jobRes.Message = "could not deserialize A"
		return &jobRes
	}

	//JSON by default does floats  ref https://javascript.info/number. Hence cannot directly convert to int
	a, ok := aRaw.(float64)
	if !ok {
		jobRes.Success = false
		jobRes.Message = "could not cast A"
		return &jobRes
	}

	bRaw, ok := args["B"]
	if !ok {
		jobRes.Success = false
		jobRes.Message = "could not deserialize B"
		return &jobRes
	}

	b, ok := bRaw.(float64)
	if !ok {
		jobRes.Success = false
		jobRes.Message = "could not cast B"
		return &jobRes
	}

	/**
	NOTE : For most cases,  It is recommended to store results in a common DB rather then use jobRes.Result. This is
	because the Jobs framework uses Redis (with a TTL) to store results , in RAM. Using Redis in this way has durability
	and cost challenges.

	For simple cases, once can use Redis as the result store

	*/

	var output int
	if args["op"] == AsyncMultiply {
		output = int(a * b)
	} else if args["op"] == AsyncAdd {
		output = int(a + b)
	}
	jobRes.Success = true
	jobRes.Result = output

	return &jobRes
}
