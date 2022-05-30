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

func TestWorkerPoolWithContext(t *testing.T) {
	sp := concurrency.NewSimplePool(maxConcurrencyTask)

	now := time.Now()
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	// this will stream every task from the batch of task we created in testTasks()
	// Since wp.tasks is buffered channel, will block writes when it's full, so
	// Submit need to take some out to make more room and unblock the pusher goroutine
	go sp.GenerateFrom(testTasks())

	// Submit will execute the tasks we have in our channel and make more room to unblock the task channel
	go sp.SubmitWithContext(ctx)

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

func TestWorkerPoolWithContext_Cancelled(t *testing.T) {
	sp := concurrency.NewSimplePool(maxConcurrencyTask)

	now := time.Now()
	ctx, cancel := context.WithCancel(context.TODO())
	go sp.SubmitWithContext(ctx)
	cancel()

	for {
		select {
		case <-sp.Done:
			fmt.Printf("we executed %v tasks in %v \n", sp.GetTaskExec(), time.Since(now))
			return
		default:
		}
	}
}

func TestWorkerPoolWithContext_Timeout(t *testing.T) {
	sp := concurrency.NewSimplePool(maxConcurrencyTask)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	now := time.Now()
	go sp.GenerateFrom(testTasks())
	go sp.SubmitWithContext(ctx)

	for {
		select {
		case <-sp.Done:
			fmt.Printf("we executed %v tasks in %v \n", sp.GetTaskExec(), time.Since(now))
			return
		default:
		}
	}
}
