package workerx

import (
	"sync"

	"github.com/cespare/xxhash/v2"
)

type Task struct {
	Entity func(...interface{})
	Args   []interface{}
}

type Workers struct {
	initializer    sync.WaitGroup
	uninitializer  sync.WaitGroup
	maxWorkerCount int
	taskQueue      []chan *Task
	signal         chan struct{}
}

func NewWorkerPool(maxWorkerCount int) *Workers {
	if maxWorkerCount < 1 {
		maxWorkerCount = 1
	}

	workers := Workers{
		initializer:    sync.WaitGroup{},
		uninitializer:  sync.WaitGroup{},
		maxWorkerCount: maxWorkerCount,
		taskQueue:      make([]chan *Task, maxWorkerCount),
		signal:         make(chan struct{}),
	}

	for idx := 0; idx < maxWorkerCount; idx++ {
		workers.taskQueue[idx] = make(chan *Task, 1024)

		workers.initializer.Add(1)
		workers.uninitializer.Add(1)

		go workers.start(idx, workers.taskQueue[idx])
	}

	// wait all goroutines to be initialized
	workers.initializer.Wait()

	return &workers
}

func (w *Workers) start(workerID int, taskChan chan *Task) {
	go func() {
		defer w.uninitializer.Done()

		w.initializer.Done()
		for {
			select {
			case <-w.signal:
				return
			case task, ok := <-taskChan:
				if !ok {
					break
				}

				// execute the specific task
				task.Entity(task.Args...)
			}
		}

	}()
}

func (w *Workers) Submit(uid string, task *Task) {
	// assign a dedicated worker to proceed the task of specific UID
	idx := xxhash.Sum64([]byte(uid)) % uint64(w.maxWorkerCount)
	if task != nil {
		w.taskQueue[idx] <- task
	}
}

func (w *Workers) Close() {
	close(w.signal)

	w.uninitializer.Wait()
}
