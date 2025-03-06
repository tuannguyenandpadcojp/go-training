package worker

import "context"

type JobHandler func(ctx context.Context) Result

type Job struct {
	ID      int
	payload string
	handler JobHandler
}
