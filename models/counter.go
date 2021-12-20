package models

import (
	"sync"
)

type Counter struct {
	Val     int
	Mtx     sync.Mutex
	Success int
	Failure int
}

func (counter *Counter) Increment() int {
	counter.Mtx.Lock()
	counter.Val++
	counter.Mtx.Unlock()
	return counter.Val
}

func (counter *Counter) HandleSuccess() int {
	counter.Mtx.Lock()
	counter.Success++
	counter.Mtx.Unlock()
	return counter.Success
}

func (counter *Counter) HandleFailure() int {
	counter.Mtx.Lock()
	counter.Failure++
	counter.Mtx.Unlock()
	return counter.Failure
}
