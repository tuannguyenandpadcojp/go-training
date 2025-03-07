package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"runtime"
)

func worker(ctx context.Context, workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
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
			results <- job.handler(ctx)
		}
	}
}

func StartWorkerPool(maxJobs, numWorkers, numJobs int) {
	pool := NewWorkerPool(maxJobs, numWorkers)
	pool.Start(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Submit jobs
	go func() {
		for i := range numJobs {
			go func() {
				var jobHandler = func(ctx context.Context) Result {
					return Result{jobID: 1, state: 1}
				}
				job := Job{ID: i, payload: fmt.Sprintf("Job %d", i), handler: jobHandler}
				if err := pool.Submit(job); err != nil {
					log.Printf("Failed to submit job with id: %d", job.ID)
				}
			}()
		}
	}()

	<-c
	log.Println("Total Goroutine before release the pool: ", runtime.NumGoroutine())
	pool.Release()
	log.Printf("Jobs success:%d - failed:%d", pool.totalSucceed, pool.totalFailed)
	time.Sleep(5 * time.Second)
	log.Println("Total Goroutine after release the pool: ", runtime.NumGoroutine())
}
