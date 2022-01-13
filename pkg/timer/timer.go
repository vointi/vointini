package timer

import (
	"time"
)

type State uint8

const (
	Unknown State = iota
	Running
	Stopped
)

func (ts State) String() string {
	switch ts {
	case Running:
		return `Running`
	case Stopped:
		return `Stopped`
	case Unknown:
		return `Unknown`
	default:
		return `?`
	}
}

type Timer struct {
	title        string
	origDuration time.Duration
	dur          time.Duration
	state        State
	startedAt    time.Time
}

func New(title string, dur time.Duration) (t *Timer) {
	t = &Timer{
		title:        title,
		origDuration: dur,
		dur:          dur,
		state:        Running,
		startedAt:    time.Now(),
	}

	go func(tmr *Timer) {
		for range time.Tick(time.Second) {
			if tmr == nil {
				break
			}

			if tmr.state == Stopped {
				continue
			}

			tmr.dur -= time.Second * 1
		}
	}(t)

	return t
}

func (t *Timer) Stop() {
	t.state = Stopped
}

func (t *Timer) Run() {
	t.state = Running
}

func (t *Timer) Get() time.Duration {
	return t.dur
}
func (t *Timer) GetTitle() string {
	return t.title
}

func (t *Timer) GetState() State {
	return t.state
}
