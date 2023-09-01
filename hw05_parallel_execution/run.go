package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Worker(wg *sync.WaitGroup, taskChan <-chan Task, m int64, count *int64) {
	defer wg.Done()
	for task := range taskChan {
		if m > 0 && atomic.LoadInt64(count) >= m {
			break
		}
		err := task()
		if err != nil {
			atomic.AddInt64(count, 1)
		}
	}
}

func Run(tasks []Task, n, m int) error {
	var errorsCount int64

	taskChan := make(chan Task)

	go func() {
		defer close(taskChan)
		for _, t := range tasks {
			if m > 0 && atomic.LoadInt64(&errorsCount) >= int64(m) {
				break
			}
			taskChan <- t
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go Worker(&wg, taskChan, int64(m), &errorsCount)
	}

	wg.Wait()

	if m > 0 && atomic.LoadInt64(&errorsCount) >= int64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
