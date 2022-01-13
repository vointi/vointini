package postgres

// Internal types to be converted to API versions

import (
	"time"
)

type entryId struct {
	Id int `db:"id"`
}

type entry struct {
	Id           int       `db:"id"`
	DateTime     time.Time `db:"fromtime"`
	ActivityName string    `db:"activity_name"`
	Achievement  int       `db:"level_achievement"`
	Description  string    `db:"descr"`
}

type entrysLevel struct {
	Level int `db:"level"`
	Id    int `db:"levelid"`
}

type entryLevel struct {
	Id               int       `db:"id"`
	Name             string    `db:"name"`
	ShowByDefault    bool      `db:"default_show"`
	DescriptionWorst string    `db:"worst_descr"`
	ShortName        string    `db:"shortname"`
	FetchPrevious    bool      `db:"get_previous"`
	AddedAt          time.Time `db:"added_at"`
}

// Tag assigned for entry
type entrysTags struct {
	TagId int `db:"tagid"`
}

type entryPrevious struct {
	Id int `db:"id"`
}

type task struct {
	Id                         int        `db:"id"`
	AddedAt                    time.Time  `db:"added_at"`
	CompletedAt                *time.Time `db:"completed_at"`
	Title                      string     `db:"title"`
	Description                string     `db:"descr"`
	Priority                   int        `db:"priority"`
	ReoccurringTaskReferenceId *int       `db:"refid"`
}

type reoccurringtask struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Title   string    `db:"title"`
}

type weight struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Weight  float32   `db:"value"`
}

type height struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Height  float32   `db:"value"`
}

type tag struct {
	Id        int       `db:"id"`
	AddedAt   time.Time `db:"added_at"`
	Name      string    `db:"name"`
	ShortName string    `db:"shortname"`
}
