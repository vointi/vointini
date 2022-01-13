package serviceitems

import "time"

// Public API formats

type EntryUpdate struct {
	Activity         string
	Description      string
	DateTime         time.Time // Precision: minutes
	LevelAchievement int
	Levels           map[int]int
	Tags             []int
}

type Entry struct {
	Id               int
	DateTime         time.Time
	ActivityName     string
	Description      string
	LevelAchievement int
	Levels           map[int]int
	Tags             []int
}

type EntryLevel struct {
	Id            int
	Name          string
	ShowByDefault bool
	ShortName     string
	Worst         string
	Previous      bool
	AddedAt       time.Time
}

type EntryLevelUpdate struct {
	Id            int
	Name          string
	ShowByDefault bool
	ShortName     string
	Worst         string
	Previous      bool
}
