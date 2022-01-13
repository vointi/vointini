package serviceitems

import "time"

type WeightAdd struct {
	Weight float32
}

type Weight struct {
	Weight  float32
	AddedAt time.Time
	Id      int
}
