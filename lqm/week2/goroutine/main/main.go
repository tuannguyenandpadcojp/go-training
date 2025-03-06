package main

import (
	"github.com/tuannguyenandpadcojp/go-training/lqm/week2/goroutine/worker"
)

func main() {
	const numJobs = 100
	const numWorkers = 20
	const maxJobs = 10
	worker.StartWorkerPool(maxJobs, numWorkers, numJobs)
}
