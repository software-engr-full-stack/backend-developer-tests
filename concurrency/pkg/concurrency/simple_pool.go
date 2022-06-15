// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"sync"
)

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())

	// Notes: I had to add this function to the interface in order for this to work.
	// If we want to keep the SimplePool interface the same, one option
	// is for NewSimplePool to return *SimplePoolRunner instead of SimplePool.
	Wait()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	var wg sync.WaitGroup

	fq := make(chan func())

	// Notes: create maxConcurrent goroutines
	for i := 0; i < maxConcurrent; i++ {
		wg.Add(1)
		go func() {
			// Notes: unbuffered channel (see declaration), will block when empty
			for fun := range fq {
				fun()
			}
			defer wg.Done()
		}()
	}

	return &SimplePoolRunner{funQueue: fq, wg: &wg}
}

type SimplePoolRunner struct {
	funQueue chan func()
	wg *sync.WaitGroup
}

func (spr *SimplePoolRunner) Submit(fun func()) {
	spr.funQueue <- fun
}

func (spr *SimplePoolRunner) Wait() {
	// Notes: originally, the "spr.wg.Wait()" line was ahead of the "close(spr.funQueue)" line.
	// This caused the program to crash with the error message "fatal error: all goroutines are asleep - deadlock!".
	// The solution was to close the channel first after all the inputs have been shoved into the channel,
	// then wait for the goroutines to finish executing.
	close(spr.funQueue)
	spr.wg.Wait()
}
