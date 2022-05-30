package main

import (
	"fmt"
	"github.com/stackpath/backend-developer-tests/concurrency"
	"sync"
	"time"
)

const (
	numberOfTasksToSend    = 50
	maxConcurrentTask      = 5
	secondsToFinishTheTask = 500 * time.Millisecond
)

func main() {
	fmt.Println("SP// Backend Developer Test - concurrency")
	var wg sync.WaitGroup
	now := time.Now()
	pool := concurrency.NewSimplePool(maxConcurrentTask) // we allow  concurrent task
	wg.Add(numberOfTasksToSend)                          // we should finish to run all task in 10sec.
	for i := 1; i <= numberOfTasksToSend; i++ {
		go pool.Submit(func() {
			fmt.Printf("Starting a task \n")
			time.Sleep(secondsToFinishTheTask) // the task takes 2sec to finish
			<-pool.Slots                       // we release the channel when the task finish
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Printf("We finish all task in:%v \n", time.Since(now))
}
