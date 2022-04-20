package postgres

// Internal types to be converted to API versions

import (
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
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

func (s entry) ConvertToAPI() *serviceitems.Entry {
	return &serviceitems.Entry{
		Id:               s.Id,
		DateTime:         s.DateTime,
		ActivityName:     s.ActivityName,
		Description:      s.Description,
		LevelAchievement: s.Achievement,
	}
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

func (s entryLevel) ConvertToAPI() *serviceitems.EntryLevel {
	return &serviceitems.EntryLevel{
		Id:            s.Id,
		Name:          s.Name,
		ShowByDefault: s.ShowByDefault,
		ShortName:     s.ShortName,
		Worst:         s.DescriptionWorst,
		Previous:      s.FetchPrevious,
		AddedAt:       s.AddedAt,
	}
}

// Tag assigned for entry
type entrysTags struct {
	TagId int `db:"tagid"`
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

func (s task) ConvertToAPI() *serviceitems.Task {
	return &serviceitems.Task{
		Id:                         s.Id,
		AddedAt:                    s.AddedAt,
		CompletedAt:                s.CompletedAt,
		Priority:                   s.Priority,
		Title:                      s.Title,
		Description:                s.Description,
		ReoccurringTaskReferenceId: s.ReoccurringTaskReferenceId,
	}
}

type reoccurringtask struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Title   string    `db:"title"`
}

func (s reoccurringtask) ConvertToAPI() *serviceitems.ReoccurringTask {
	return &serviceitems.ReoccurringTask{
		Id:      s.Id,
		AddedAt: s.AddedAt,
		Title:   s.Title,
	}
}

// weight has person's weight
type weight struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Weight  float32   `db:"value"`
}

func (s weight) ConvertToAPI() *serviceitems.Weight {
	return &serviceitems.Weight{
		Id:      s.Id,
		AddedAt: s.AddedAt,
		Weight:  s.Weight,
	}
}

// height has person's height
type height struct {
	Id      int       `db:"id"`
	AddedAt time.Time `db:"added_at"`
	Height  float32   `db:"value"`
}

func (s height) ConvertToAPI() *serviceitems.Height {
	return &serviceitems.Height{
		Id:      s.Id,
		AddedAt: s.AddedAt,
		Height:  s.Height,
	}
}

type tag struct {
	Id        int       `db:"id"`
	AddedAt   time.Time `db:"added_at"`
	Name      string    `db:"name"`
	ShortName string    `db:"shortname"`
}

func (s tag) ConvertToAPI() *serviceitems.Tag {
	return &serviceitems.Tag{
		Id:        s.Id,
		AddedAt:   s.AddedAt,
		Name:      s.Name,
		ShortName: s.ShortName,
	}
}

type resolution struct {
	Id           int        `db:"id"`
	Name         string     `db:"name"`
	DecisionDate *time.Time `db:"decisiondate"`
	SentDate     *time.Time `db:"sentdate"`
	StartDate    time.Time  `db:"startdate"`
	EndDate      *time.Time `db:"enddate"`
	EntityId     int        `db:"entityid"`
	AddedAt      time.Time  `db:"added_at"`
}

func (s resolution) ConvertToAPI() *serviceitems.Resolution {
	return &serviceitems.Resolution{
		Id:           s.Id,
		Name:         s.Name,
		DecisionDate: s.DecisionDate,
		SentDate:     s.SentDate,
		StartDate:    s.StartDate,
		EndDate:      s.EndDate,
		EntityId:     s.EntityId,
		AddedAt:      s.AddedAt,
	}
}

type resolutionFile struct {
	Id           int       `db:"id"`
	ResolutionId int       `db:"resolutionid"`
	AddedAt      time.Time `db:"added_at"`
	Filename     string    `db:"filename"`
	ContentType  string    `db:"ctype"`
}

func (s resolutionFile) ConvertToAPI() *serviceitems.ResolutionFile {
	return &serviceitems.ResolutionFile{
		Id:           s.Id,
		ResolutionId: s.ResolutionId,
		AddedAt:      s.AddedAt,
		Filename:     s.Filename,
		ContentType:  s.ContentType,
	}
}

type resolutionEntity struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (s resolutionEntity) ConvertToAPI() *serviceitems.ResolutionEntity {
	return &serviceitems.ResolutionEntity{
		Id:   s.Id,
		Name: s.Name,
	}
}
