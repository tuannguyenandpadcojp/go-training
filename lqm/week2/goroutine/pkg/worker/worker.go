package worker

import (
	"context"
	"log"
	"sync"
)

type Result struct {
	JobID int
	State int // 0: Failed, 1: Success
}

type JobHandler func(ctx context.Context) Result

type Job struct {
	ID      int
	Payload string
	Handler JobHandler
}

func Worker(ctx context.Context, workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				// Job channels are closed
				return
			}
			log.Printf("Worker %d processed job %d", workerID, job.ID)
			results <- job.Handler(ctx)
		}
	}
}
