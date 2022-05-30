// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stackpath/backend-developer-tests/concurrency"
)

const (
	tasksCount         = 30
	maxConcurrencyTask = 2
)

func TestWorkerPoolWithoutContext(t *testing.T) {
	sp := concurrency.NewSimplePool(maxConcurrencyTask)

	now := time.Now()
	// this will stream every task from the batch of task we created in testTasks()
	// Since wp.tasks is buffered channel, will block writes when it's full, so
	// Submit need to take some out to make more room and unblock the pusher goroutine
	go sp.GenerateFrom(testTasks())
	// Submit will execute the tasks we have in our channel and make more room to unblock the task channel
	go sp.Submit()

	for {
		select {
		case <-sp.Done:
			// we should execute the task in (tasksCount / maxConcurrencyTask)
			// i.e. 20 tasks to execute, and we allow just 5 concurrent task
			// the time should be is: 4s. 20/5 = 4
			// i.e. 30 tasks to execute, and we allow just 2 concurrent task
			// the time should be is : 15s. 30/2 = 15
			fmt.Printf("we executed %v tasks in %v \n", sp.GetTaskExec(), time.Since(now))
			return
		default:
		}
	}
}

func testTasks() []func(int64, context.Context) {
	tasks := make([]func(int64, context.Context), tasksCount)
	for i := 0; i < tasksCount; i++ {
		tasks[i] = func(taskNumber int64, ctx context.Context) {
			fmt.Printf("Executing the task number: %v \n", taskNumber)
			// each task will take 1 seconds to finish (super slow task)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finish the task number: %v \n", taskNumber)
		}
	}
	return tasks
}
