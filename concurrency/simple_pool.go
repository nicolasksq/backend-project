// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

// SimplePoolSubmitter is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePoolSubmitter interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(task func())
}

type SimplePool struct {
	workersCount int
	Slots        chan func()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	return SimplePool{
		workersCount: maxConcurrent,
		Slots:        make(chan func(), maxConcurrent),
	}
}

func (s SimplePool) Submit(task func()) {
	s.Slots <- task
	go task() // we execute the task asynchronously
}
