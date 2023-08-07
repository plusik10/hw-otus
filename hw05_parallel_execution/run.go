package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	workersPool := make(chan struct{}, n)
	maxErrors := int64(m)
	var errCount int64

	for i, task := range tasks {
		if atomic.LoadInt64(&errCount) >= maxErrors && maxErrors > 0 {
			wg.Add(-(len(tasks) - i))
			return ErrErrorsLimitExceeded
		}
		workersPool <- struct{}{}
		go func(task Task) {
			defer func() {
				<-workersPool
				wg.Done()
			}()
			taskError := task()
			if taskError != nil {
				atomic.AddInt64(&errCount, 1)
			}
		}(task)
	}
	wg.Wait()
	return nil
}
