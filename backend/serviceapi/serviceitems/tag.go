package serviceitems

import "time"

type Tag struct {
	Id        int
	AddedAt   time.Time
	Name      string
	ShortName string // Machine-parsable name
}

type TagUpdate struct {
	Id        int
	Name      string
	ShortName string // Machine-parsable name
}
