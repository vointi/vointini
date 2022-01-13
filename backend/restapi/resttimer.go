package restapi

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/vointini/vointini/pkg/timer"
	"io"
	"net/http"
	"strconv"
	"time"
)

var counter = uint64(0)

// timerAdd adds a new Timer which counts down
func (restapi *restAPI) timerAdd(w http.ResponseWriter, r *http.Request) {
	var item DTOTimerAdd
	if err := readStruct(r.Body, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	if item.Seconds == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c := counter
	counter++

	// Add new Timer
	restapi.timers[c] = timer.New(item.Title, time.Second*time.Duration(item.Seconds))

	go func(tid uint64) {
		for range time.Tick(time.Second) {
			tmr, ok := restapi.timers[tid]
			if !ok {
				// Timer doesn't exist
				break
			}

			dur := tmr.Get()

			itm := DTOTimer{
				Id:        tid,
				Seconds:   dur.Seconds(),
				Formatted: dur.String(),
			}

			b, err := json.Marshal(itm)
			if err != nil {
				panic(err)
				return
			}

			// Send SSE event JSON
			restapi.sseServer.SendMessage(
				`/api/v1/timers/events`,
				sse.SimpleMessage(string(b)),
			)
		}

	}(c)

	_, _ = io.WriteString(w,
		fmt.Sprintf(`{"id": %d}`, c),
	)
}

// timerStop stops a Timer with a given ID
func (restapi *restAPI) timerStop(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, `id`), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	if _, ok := restapi.timers[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	restapi.timers[id].Stop()

	b, err := json.Marshal(&DTOOK{
		Msg: `ok`,
	})
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// timerRemove removes a Timer with a given ID
func (restapi *restAPI) timerRemove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, `id`), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if _, ok := restapi.timers[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	restapi.timers[id] = nil

	b, err := json.Marshal(&DTOOK{
		Msg: `ok`,
	})
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// timerList lists timers which exists
func (restapi *restAPI) timerList(w http.ResponseWriter, r *http.Request) {
	var ids []uint64
	for i, t := range restapi.timers {
		if t == nil {
			continue
		}

		ids = append(ids, i)
	}

	if len(ids) == 0 {
		// no timers
		_, _ = io.WriteString(w, `[]`)
		return
	}

	b, err := json.Marshal(&ids)
	if err != nil {
		panic(err)
		return
	}

	_, _ = w.Write(b)
}
