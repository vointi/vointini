package restapi

import (
	"encoding/json"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vointini/vointini/backend/restapi/locales"
	"github.com/vointini/vointini/backend/serviceapi"
	"github.com/vointini/vointini/pkg/timer"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"net/http"
	"strconv"
)

// New creates a new router which serves REST API
func New(service *serviceapi.Service, defaultLanguage language.Tag) (router *chi.Mux) {
	endpoint := newApi(service, defaultLanguage)

	router = chi.NewRouter()
	router.Use(middleware.SetHeader(`Content-Type`, `application/json; charset=utf-8`))

	// Entries
	router.Get(`/entries/{year}/{month}/{day}/{hour}/{minute}`, endpoint.entriesMinute)
	router.Get(`/entries/{year}/{month}/{day}`, endpoint.entriesDay) // List
	router.Post(`/entries/{year}/{month}/{day}/{hour}/{minute}`, endpoint.entryUpdate)
	router.Get(`/entries/levels`, endpoint.entriesLevels) // List
	router.Post(`/entries/level/{id}`, endpoint.entryLevelUpdate)
	router.Post(`/entries/level`, endpoint.entryLevelUpdate)

	// Tags
	router.Get(`/tags`, endpoint.tagList)        // List
	router.Post(`/tag/{id}`, endpoint.tagUpdate) // Update existing
	router.Get(`/tag/{id}`, endpoint.tagGet)

	// Tasks
	router.Get(`/tasks`, endpoint.taskList) // List
	router.Get(`/task/{id}`, endpoint.taskGet)
	router.Post(`/task/{id}`, endpoint.taskUpdate) // Update existing
	router.Post(`/task`, endpoint.taskUpdate)      // New

	// Re-occurring tasks
	router.Get(`/reoccurring-tasks`, endpoint.reoccurringTaskList) // List
	router.Get(`/reoccurring-task/{id}`, endpoint.reoccurringTaskGet)
	router.Post(`/reoccurring-task/{id}`, endpoint.reoccurringTaskUpdate) // Update existing
	router.Post(`/reoccurring-task`, endpoint.reoccurringTaskUpdate)      // New

	// Resolutions
	router.Get(`/resolutions`, endpoint.resolutionsList)        // List
	router.Post(`/resolution/{id}`, endpoint.resolutionsUpdate) // Update existing
	router.Get(`/resolution/{id}`, endpoint.resolutionsGet)

	router.Get(`/resolution-entities`, endpoint.resolutionsEntityList)  // List entities
	router.Get(`/resolution-files`, endpoint.resolutionFileList)        // List files
	router.Get(`/resolution-files/{id}`, endpoint.resolutionFileList)   // List files of a given resolution
	router.Post(`/resolution-file/{id}`, endpoint.resolutionFileUpload) // Upload a file to a given resolution

	// Weight
	router.Get(`/weight`, endpoint.weightList)    // List
	router.Post(`/weight`, endpoint.weightUpdate) // New

	// Height
	router.Get(`/height`, endpoint.heightList)    // List
	router.Post(`/height`, endpoint.heightUpdate) // New

	// Test: MADRS
	router.Get(`/tests/madrs`, endpoint.testMADRSResults) // List
	router.Post(`/tests/madrs`, endpoint.testMADRS)       // New

	// Timers

	router.Post(`/timer`, endpoint.timerAdd)  // New
	router.Get(`/timers`, endpoint.timerList) // List
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
	tr        *message.Printer
}

func newApi(api *serviceapi.Service, defaultLanguage language.Tag) *restAPI {
	return &restAPI{
		api:    api,
		timers: make(map[uint64]*timer.Timer),
		tr: message.NewPrinter(defaultLanguage,
			message.Catalog(locales.Translations),
		),
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
