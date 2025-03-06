package worker

import (
	"context"
	"log"
	"sync"
	"fmt"
)

type Pool struct {
	// Mutex to protect the pool state
	mutex sync.Mutex

	// Context and cancel function to notify all workers to stop
	ctx        context.Context
	cancelFunc context.CancelFunc

	// Defines the number of workers and job queue
	workers int
	jobs chan Job

	// Wait group to wait for all workers to finish
	workerWaitGroup sync.WaitGroup

	// Wait group to wait for the goroutine aggregate the job result
	resultWaitGroup sync.WaitGroup

	// Pool state
	running bool

	// Non-blocking mode
	nonBlocking bool

	// Job result
	results      chan Result
	totalSucceed int
	totalFailed  int
}

type PoolOpt func(p *Pool)

func WithNonBlocking(p *Pool) {
	p.nonBlocking = true
}

func NewWorkerPool(maxJobs, numWorkers int, opts ...PoolOpt) *Pool {
	p := &Pool{
		workers: numWorkers,
		jobs: make(chan Job, maxJobs),
		results: make(chan Result, maxJobs),
		workerWaitGroup: sync.WaitGroup{},
		resultWaitGroup: sync.WaitGroup{},
		mutex: sync.Mutex{},
	}
	for _, opt := range(opts) {
		opt(p)
	}
	return p
}

func (p *Pool) Start(ctx context.Context) {
	p.mutex.Lock()

	// Unlock after done 
	defer p.mutex.Unlock()

	if p.running {
		log.Println("Worker Pool is running")
		return
	}

	p.running = true
	p.ctx, p.cancelFunc = context.WithCancel(ctx)
	p.workerWaitGroup.Add(p.workers)
	
	// spawn worker goroutine
	for i := range p.workers {
		go worker(p.ctx, i, p.jobs, p.results, &p.workerWaitGroup)
	}

	// aggregate job's result
	p.resultWaitGroup.Add(1)
	go func() {
		defer p.resultWaitGroup.Done()
		for result := range p.results {
			if result.state == 1 {
				p.totalSucceed++
			} else {
				p.totalFailed++
			}
		}
	}()
}

func (p *Pool) Release() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// close the Jobs channel to prevent dispatcher send jobs
	close(p.jobs)
	// wait for all workers to finish processing the rest of jobs
	p.workerWaitGroup.Wait()
	// call context.CancelFunc to notify all workers to stop
	p.cancelFunc()
	// close the result channel to stop the goroutine aggregate the job's result
	close(p.results)
	// wait for the goroutine aggregate the job result is done
	p.resultWaitGroup.Wait()
	p.running = false
}

// Submit is a dispatcher feeds jobs for the worker
func (p *Pool) Submit(job Job) error {
	if !p.running {
		return fmt.Errorf("pool is closed")
	}
	if p.nonBlocking {
		select {
		case p.jobs <- job:
			// If the job channel has space, the job is sent and the function returns nil.
			return nil
		default:
			// If the job channel is full, the default case is executed and an error is returned.
			return fmt.Errorf("job queue is full")
		}
	} else {
		// blocking if the jobs channel is full
		p.jobs <- job
		return nil
	}
}