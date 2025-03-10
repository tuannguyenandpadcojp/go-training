package main

import (
	"context"
	"testing"

	"github.com/tuannguyenandpadcojp/go-training/lqm/utils"
	"github.com/tuannguyenandpadcojp/go-training/lqm/week2/goroutine/pkg/worker"
	"go.uber.org/goleak"
)

func TestWorkerPool(t *testing.T) {
	pool, err := worker.NewWorkerPool(3, 10)
	utils.AssertError(t, err, false)
	pool.Start(context.Background())

	var jobHandlerSuccess = func(ctx context.Context) worker.Result {
		return worker.Result{JobID: 1, State: 1}
	}

	var jobHandlerFail = func(ctx context.Context) worker.Result {
		return worker.Result{JobID: 0, State: 0}
	}

	var mockJobs []worker.Job
	for i := range 5 {
		mockJobs = append(mockJobs, worker.Job{ID: i, Payload: "mock", Handler: jobHandlerSuccess}) // Success
	}
	for i := 5; i < 10; i++ {
		mockJobs = append(mockJobs, worker.Job{ID: i, Payload: "mock", Handler: jobHandlerFail}) // Fail
	}

	for _, job := range mockJobs {
		if err := pool.Submit(job); err != nil {
			t.Errorf("Error submiting job id %d", job.ID)
		}
	}

	pool.Release()

	expectSuccess, expectFail := 5, 5
	resultSuccess, resultFail := pool.TotalSucceed, pool.TotalFailed

	utils.AssertCorrectResult(t, resultSuccess, expectSuccess)
	utils.AssertCorrectResult(t, resultFail, expectFail)

	goleak.VerifyNone(t)
}

func TestWorkerPoolNonBlocking(t *testing.T) {
	pool, err := worker.NewWorkerPool(5, 3, worker.WithNonBlocking)
	utils.AssertError(t, err, false)

	pool.Start(context.Background())

	wait := make(chan struct{})
	var blockingHandler = func(ctx context.Context) worker.Result {
		<-wait
		return worker.Result{JobID: 1, State: 1}
	}

	for i := range 10 {
		job := worker.Job{ID: i, Payload: "mock"}
		if i < 5 {
			job.Handler = blockingHandler
		}
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
	if pool.TotalSucceed != 5 {
		t.Errorf("Expected 5 successful jobs, got %d", pool.TotalSucceed)
	}
	if pool.TotalFailed != 0 {
		t.Errorf("Expected 0 failed jobs, got %d", pool.TotalFailed)
	}

	goleak.VerifyNone(t)
}

func TestWorkerPoolInvalidInput(t *testing.T) {
	inputTests := []struct {
		maxJobs     int
		numWorkers  int
		shouldError bool
	}{
		{maxJobs: 1, numWorkers: 1, shouldError: false},
		{maxJobs: -1, numWorkers: 1, shouldError: true},
		{maxJobs: 2, numWorkers: -1, shouldError: true},
	}

	for _, input := range inputTests {
		_, result := worker.NewWorkerPool(input.maxJobs, input.numWorkers)
		utils.AssertError(t, result, input.shouldError)
	}
}
