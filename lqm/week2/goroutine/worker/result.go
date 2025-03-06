package worker

type Result struct {
	jobID int
	state int // 0: Failed, 1: Success
}
