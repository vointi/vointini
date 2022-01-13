package serviceitems

import "time"

type HeightAdd struct {
	Height float32
}

type Height struct {
	Height  float32
	AddedAt time.Time
	Id      int
}
