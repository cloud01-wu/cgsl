package workerx

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	var (
		WorkerCount = 8
		TaskCount   = 64
	)

	// wait all tasks to be done
	wait := sync.WaitGroup{}

	tasks := []*Task{}

	// compose the tasks
	for idx := 0; idx < TaskCount; idx++ {
		wait.Add(1)
		entity := func(args ...interface{}) {
			taskID := args[0].(int)
			arg1 := args[1].(int)
			t.Logf("task %d was executed, arg is %d\n", taskID, arg1)
			wait.Done()
		}

		args := []interface{}{idx, time.Now().Nanosecond()}

		task := &Task{
			Entity: entity,
			Args:   args,
		}

		tasks = append(tasks, task)
	}

	// create the worker
	workers := NewWorkerPool(WorkerCount)
	defer workers.Close()

	// submit the tasks one by one
	for idx1 := 0; idx1 < TaskCount; idx1++ {
		workers.Submit(strconv.Itoa(rand.Intn(8-1)+1), tasks[idx1])
	}

	wait.Wait()
}
