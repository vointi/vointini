package restapi

import (
	"encoding/json"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vointini/vointini/backend/serviceapi"
	"github.com/vointini/vointini/pkg/timer"
	"io"
	"net/http"
	"strconv"
)

// New creates a new router which serves REST API
func New(service *serviceapi.Service) (router *chi.Mux) {
	endpoint := newApi(service)

	router = chi.NewRouter()
	router.Use(middleware.SetHeader(`Content-Type`, `application/json; charset=utf-8`))

	// Entries
	router.Get(`/entries/{year}/{month}/{day}/{hour}/{minute}`, endpoint.entriesMinute)
	router.Get(`/entries/{year}/{month}/{day}`, endpoint.entriesDay)
	router.Post(`/entries/{year}/{month}/{day}/{hour}/{minute}`, endpoint.entryUpdate)
	router.Get(`/entries/levels`, endpoint.entriesLevels)
	router.Post(`/entries/level/{id}`, endpoint.entryLevelUpdate)
	router.Post(`/entries/level`, endpoint.entryLevelUpdate)

	// Tags
	router.Get(`/tags`, endpoint.tagList)
	router.Post(`/tag/{id}`, endpoint.tagUpdate) // Update existing
	router.Get(`/tag/{id}`, endpoint.tagGet)

	// Tasks
	router.Get(`/tasks`, endpoint.taskList)
	router.Get(`/task/{id}`, endpoint.taskGet)
	router.Post(`/task/{id}`, endpoint.taskUpdate) // Update existing
	router.Post(`/task`, endpoint.taskUpdate)      // New

	// Re-occurring tasks
	router.Get(`/reoccurring-tasks`, endpoint.reoccurringTaskList)
	router.Get(`/reoccurring-task/{id}`, endpoint.reoccurringTaskGet)
	router.Post(`/reoccurring-task/{id}`, endpoint.reoccurringTaskUpdate) // Update existing
	router.Post(`/reoccurring-task`, endpoint.reoccurringTaskUpdate)      // New

	// Weight
	router.Get(`/weight`, endpoint.weightList)
	router.Post(`/weight`, endpoint.weightUpdate) // New

	// Height
	router.Get(`/height`, endpoint.heightList)
	router.Post(`/height`, endpoint.heightUpdate) // New

	// Test: MADRS
	router.Get(`/tests/madrs`, endpoint.testMADRSResults)
	router.Post(`/tests/madrs`, endpoint.testMADRS) // New

	// Timers

	router.Post(`/timer`, endpoint.timerAdd)
	router.Get(`/timers`, endpoint.timerList)
	router.Get(`/timer/{id}/stop`, endpoint.timerStop)
	router.Get(`/timer/{id}/remove`, endpoint.timerRemove)

	// SSE for Timer events
	router.Handle(`/timers/events`, endpoint.sseServer)

	return router
}

type restAPI struct {
	api       *serviceapi.Service
	timers    map[uint64]*timer.Timer
	sseServer *sse.Server
}

func newApi(api *serviceapi.Service) *restAPI {
	return &restAPI{
		api:    api,
		timers: make(map[uint64]*timer.Timer),
		sseServer: sse.NewServer(&sse.Options{
			RetryInterval: 5,
			//Logger:        log.New(os.Stdout, `SSE: `, 0),
			Logger: nil,
		}),
	}
}

func readStruct(r io.Reader, i interface{}) (err error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func getIntParam(req *http.Request, name string) (val int, err error) {
	tmp := int64(-1)

	intStr := chi.URLParam(req, name)
	if intStr != `` {
		tmp, err = strconv.ParseInt(intStr, 10, 64)
	}

	return int(tmp), err
}
