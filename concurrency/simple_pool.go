// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"context"
	"fmt"
	"sync"
)

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
	SubmitWithContext(func(context.Context))
}

type WorkerPool struct {
	workersCount int
	tasks        chan func(int64, context.Context)
	Done         chan struct{}

	mu            sync.Mutex
	taskExecution int64
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) WorkerPool {
	return WorkerPool{
		workersCount: maxConcurrent,
		tasks:        make(chan func(int64, context.Context), maxConcurrent),
		Done:         make(chan struct{}),
	}
}

// Submit will throw goroutines which will execute the function in tasks channel async.
func (wp *WorkerPool) Submit() {
	var wg sync.WaitGroup
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go wp.exec(context.TODO(), &wg)
	}
	wg.Wait()
	// if the goroutines have finished and there are no more tasks to execute, we will close the Done channel.
	close(wp.Done)
}

func (wp *WorkerPool) SubmitWithContext(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go wp.exec(ctx, &wg)
	}
	wg.Wait()
	close(wp.Done)
}

// GenerateFrom will populate the channel with the given function
// (int64 has been passed int the func - just to see how many task we are executing)
func (wp *WorkerPool) GenerateFrom(tasksBulk []func(int64, context.Context)) {
	for i := range tasksBulk {
		wp.tasks <- tasksBulk[i]
	}
	close(wp.tasks)
}

// add task has been created just to see how many times we are executing the task
// we increased concurrently safely +1
func (wp *WorkerPool) addTaskExec() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	wp.taskExecution++
}

// GetTaskExec get the number of tasks executed.
func (wp *WorkerPool) GetTaskExec() int64 {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	return wp.taskExecution
}

// exec will execute the function we have in job channel.
// It will read from wp.tasks until the context is cancelled or there is no more tasks to run
// also we are summing by one (wp.addTaskExec()) our counter to see how many task has been executed.
func (wp *WorkerPool) exec(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok {
				return
			}
			wp.addTaskExec()
			task(wp.taskExecution, ctx)
		case <-ctx.Done():
			fmt.Printf("cancelled execution. Error detail: %v\n", ctx.Err())
			return
		}
	}
}
