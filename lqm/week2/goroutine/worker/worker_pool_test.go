package worker

import (
	"context"
	"testing"

	"github.com/tuannguyenandpadcojp/go-training/lqm/utils"
	"go.uber.org/goleak"
)

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(3, 10)
	pool.Start(context.Background())

	var jobHandlerSuccess = func (ctx context.Context) Result {
		return Result{jobID: 1, state: 1}
	}

	var jobHandlerFail = func (ctx context.Context) Result {
		return Result{jobID: 0, state: 0}
	}

	var mockJobs []Job
	for i := range 5 {
		mockJobs = append(mockJobs, Job{ID: i, payload: "mock", handler: jobHandlerSuccess}) // Success 
	}
	for i := 5; i < 10; i++ {
		mockJobs = append(mockJobs, Job{ID: i, payload: "mock", handler: jobHandlerFail}) // Fail
	}

	for _, job := range mockJobs {
		if err := pool.Submit(job); err != nil {
			t.Errorf("Error submiting job id %d", job.ID)
		}
	}

	pool.Release()

	expectSuccess, expectFail := 5, 5
	resultSuccess, resultFail := pool.totalSucceed, pool.totalFailed

	utils.AssertCorrectResult(t, resultSuccess, expectSuccess)
	utils.AssertCorrectResult(t, resultFail, expectFail)

	goleak.VerifyNone(t)
}

func TestWorkerPoolNonBlocking(t *testing.T) {
	pool := NewWorkerPool(3, 5, WithNonBlocking)
	pool.Start(context.Background())

	wait := make(chan struct{})
	var blockingHandler = func(ctx context.Context) Result {
		<-wait
		return Result{jobID: 1, state: 1}
	}

	for i := range 10 {
		job := Job{ID: i, payload: "mock", handler: blockingHandler}
		err := pool.Submit(job)
		if i < 5 && err != nil {
			t.Errorf("Failed to submit job: %v", err)
		}
		if i >= 5 && err == nil {
			t.Errorf("Expected job queue to be full, but job was submitted successfully")
		}
	}

	close(wait)

	// Release the pool ans closed
	pool.Release()

	// Check results
	if pool.totalSucceed != 5 {
		t.Errorf("Expected 5 successful jobs, got %d", pool.totalSucceed)
	}
	if pool.totalFailed != 0 {
		t.Errorf("Expected 0 failed jobs, got %d", pool.totalFailed)
	}

	goleak.VerifyNone(t)
}