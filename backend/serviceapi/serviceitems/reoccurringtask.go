package serviceitems

import "time"

type ReoccurringTaskUpdate struct {
	Id       int
	Title    string
	Duration time.Duration
}

type ReoccurringTask struct {
	Id       int
	AddedAt  time.Time
	Title    string
	Duration time.Duration
}
