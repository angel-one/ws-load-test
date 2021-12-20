# Handling Async APIs
This project has a jobs folder which enables easy async execution of jobs.
It uses [Asynq](https://github.com/hibiken/asynq) as the backend but this is abstracted from the rest of the app by the
``jobs`` package. The ``async_maths`` API demos how to use the Jobs package. Redis is used for both the job  queues and
the results store

## Jobs
The unit of async work is a Job. Every Job belongs to a Job Type - which is identified by an unique name.
Every job type also has a  handler where the business logic to execute async is hosted. The core interface is :
```type JobHandler func(map[string]interface{}) *JobResult```
A sample can be found in ``business\async_maths.go``  package - the ``MathsHandler`` method

Jobs are executed in Workers which are goroutines spawned by the runtime. This is where the jobs are executed. When a job 
is created, it also needs some arguments - which here is a ``map[string]interface{}``. The job handler once done returns
a `` JobStatus`` instance.

Every job instance is identified by a unique string id.

Each job instance can be in one of the following states :
- Active : Job is being worked on
- Pending : Indicates that the task is ready to be processed by Handler.
- Future : Indicates that the task is scheduled to be processed some time in the future.
- Retrying :Indicates that the task has previously failed and scheduled to be processed some time
- Failed : Indicates that the task is  failed 
- Completed :Indicates that the task is processed successfully



## Jobs Key Methods
The main "create job" APIs are :
- NewJobController : 
Creates a new instance of a Jobs Controller. Needed to be done once in the app
- AddHandler :
Each Jobs is identified by a name (string). The AddHandler method adds a handler for a specific job type  name. 
- Start :
Start the Job controller and create background workers. This is done once at ``main.go``
- CreateJob :
This is the method to create a job for every request. Here you give the JobType (string)  and the params for the instance
- GetStatus :
Gets the status of the job, and if it's complete the result is available as part of ``JobStatus`` object returned

> **_NOTE:_** Calling AddHandler after Start is a noop. Please see the ```initAsyncModules ``` in main.go for details


## Sample Requests

- Create an async maths job
```shell
curl -X POST \
  http://localhost:8080/maths \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: c305cb26-1c5e-4c64-a493-3cb756995a47' \
  -d '   
   {
     "a": 2,
     "b": 3,
     "op": "math:multiply"
   }'
```
This returns 
```json
{
    "success": true,
    "res": "default@f6404090-f2bb-4490-8478-075aa95b29d1"
}
```
here res is the id of the job

- Get the status (and result of the job)

```shell
curl -X GET \
  http://localhost:8080/maths/default@f6404090-f2bb-4490-8478-075aa95b29d1 \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 5d024699-7029-7a78-2e32-de4558e7a41c' \
  -d '   
   {
     "a": 2,
     "b": 3,
     "op": "math:multiply"
   }'
```
which returns
```json
{
    "state_code": 6,
    "status": "COMPLETED",
    "res": 6
}
```

Note : "state": 6 ; here 6 refers to job in Completed state

## Asynq Internals
- Details about Aynsq : [Wiki](https://github.com/hibiken/asynq/wiki/Getting-Started) 
- Task State Diagram : [[here](https://github.com/hibiken/asynq/wiki/Life-of-a-Task)]

