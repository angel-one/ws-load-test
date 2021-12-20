package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/hibiken/asynq"
)

type JobHandler func(map[string]interface{}) *JobResult

type JobController struct {
	delegateClient *asynq.Client
	inspector      *asynq.Inspector
	handlers       map[string]JobHandler
	redisAddr      string
	jobsTTLHrs     time.Duration
	nWorkers       int
}

type JobState int

const (
	// Indicates that the task is currently being processed by Handler.
	Active JobState = iota + 1

	// Indicates that the task is ready to be processed by Handler.
	Pending

	// Indicates that the task is scheduled to be processed some time in the future.
	Future

	// Indicates that the task has previously failed and scheduled to be processed some time in the future.
	Retrying

	// Indicates that the task is archived and stored for inspection purposes. Failure
	Failed

	// Indicates that the task is processed successfully and retained until the retention TTL expires.
	Completed
)

type JobStatus struct {
	State  JobState `json:"state_code"` // 1 is Active 6 is Completed
	Status string   `json:"status"`
	Result []byte   `json:"res"`
}

type JobResult struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Result  interface{} `json:"res"`
}

// Create a new instance of controller
func NewJobController(redisAddr string, jobsRetentionHrs uint, numberOfWorkers int) *JobController {
	// Make a asynq client
	return &JobController{
		delegateClient: asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr}),
		inspector:      asynq.NewInspector(asynq.RedisClientOpt{Addr: redisAddr}),
		redisAddr:      redisAddr,
		jobsTTLHrs:     time.Duration(jobsRetentionHrs) * time.Hour,
		nWorkers:       numberOfWorkers,
		handlers:       make(map[string]JobHandler),
	}

}

func (j *JobController) Start() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: j.redisAddr},
		asynq.Config{Concurrency: j.nWorkers},
	)

	genericHandler := func(ctx context.Context, t *asynq.Task) error {
		// all args are a map of string to interface
		var args map[string]interface{}
		if err := json.Unmarshal(t.Payload(), &args); err != nil {
			return err
		}

		// This is where actual handler gets exectued
		theHandler, ok := j.handlers[t.Type()]
		if !ok {
			return errors.New("handler not found for task :" + t.Type())
		}

		res := theHandler(args)
		payload, err := json.Marshal(res)
		if err != nil {
			return err
		}

		t.ResultWriter().Write(payload)
		return nil
	}

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(genericHandler)); err != nil {
		log.Fatal(err)
	}

}

func (j *JobController) Close() {
	j.delegateClient.Close()
}

func (j *JobController) AddHandler(jobName string, handler JobHandler) {
	j.handlers[jobName] = handler
}

func (j *JobController) CreateJob(name string, args map[string]interface{}) (string, error) {
	// Enqueue a job named  with the specified parameters.
	payload, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	task := asynq.NewTask(name, payload)
	info, err := j.delegateClient.Enqueue(task, asynq.Retention(j.jobsTTLHrs))
	if err != nil {
		log.Printf("could not schedule task: %v\n", err)
		return "", err
	}

	return getTaskId(info), nil
}

// Get status of a Job. JobState denotes the state of a task.
func (j *JobController) GetStatus(id string) (*JobStatus, error) {

	var jobStatus JobStatus
	qname, asynqId := getQueueAndTaskId(id)
	info, err := j.inspector.GetTaskInfo(qname, asynqId)
	if err != nil {
		log.Printf("could not get task status: %v\n", err)
		return &jobStatus, err
	}

	// TODO handle unknown info.state - can happen only with ver changes of asynq
	jobStatus.State, jobStatus.Status = statusTranslator[info.State].code, statusTranslator[info.State].status
	jobStatus.Result = info.Result

	return &jobStatus, nil
}

// Transalation of Asynq codes to Jobs status
var statusTranslator = map[asynq.TaskState]struct {
	status string
	code   JobState
}{
	asynq.TaskStateActive:    {"IN_WORK", Active},
	asynq.TaskStatePending:   {"TO_BE_SCHEDULED", Pending},
	asynq.TaskStateScheduled: {"FUTURE", Future},
	asynq.TaskStateRetry:     {"RETRYING", Retrying},
	asynq.TaskStateArchived:  {"FAILED", Failed},
	asynq.TaskStateCompleted: {"SUCCESS", Completed},
}

func getTaskId(info *asynq.TaskInfo) string {
	return info.Queue + "@" + info.ID
}

func getQueueAndTaskId(id string) (string, string) {
	res := strings.Split(id, "@")
	return res[0], res[1]
}
