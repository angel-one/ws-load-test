package models

import (
	"sync"
)

type Counter struct {
	Val int
	mtx sync.Mutex
	success, failure int
}

func (counter *Counter) Increment() int {
	counter.mtx.Lock()
	counter.Val++
	counter.mtx.Unlock()
	return counter.Val
}

func (counter *Counter) Success() int {
	counter.mtx.Lock()
	counter.success++
	counter.mtx.Unlock()
	return counter.success
}

func (counter *Counter) Failure() int {
	counter.mtx.Lock()
	counter.failure++
	counter.mtx.Unlock()
	return counter.failure
}
