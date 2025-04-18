package task

import "time"

type Task struct {
	des        string
	finished   bool
	createdAt  time.Time
	finishedAt *time.Time
}
