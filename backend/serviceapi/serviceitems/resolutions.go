package serviceitems

import "time"

type ResolutionsUpdate struct {
	Id           int
	Name         string
	DecisionDate *time.Time
	SentDate     *time.Time
	StartDate    time.Time
	EndDate      *time.Time
	EntityId     int
}

type Resolution struct {
	Id           int
	AddedAt      time.Time
	Name         string
	DecisionDate *time.Time
	SentDate     *time.Time
	StartDate    time.Time
	EndDate      *time.Time
	EntityId     int
}

type ResolutionEntity struct {
	Id   int
	Name string
}

type ResolutionFile struct {
	Id           int
	ResolutionId int
	AddedAt      time.Time
	Filename     string
}
