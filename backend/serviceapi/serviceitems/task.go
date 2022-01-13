package serviceitems

import "time"

type TaskSearchOption uint8

const (
	Unknown TaskSearchOption = iota
	OngoingTasks
	CompletedTasks
)

type TaskUpdate struct {
	Id                         int
	Priority                   int
	Title                      string
	Description                string
	ReoccurringTaskReferenceId *int
}

type Task struct {
	Id                         int
	AddedAt                    time.Time
	CompletedAt                *time.Time
	Priority                   int
	Title                      string
	Description                string
	ReoccurringTaskReferenceId *int
}
