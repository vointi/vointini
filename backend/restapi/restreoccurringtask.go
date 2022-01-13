package restapi

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Re-occurring task related

func (restapi restAPI) reoccurringTaskUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	id := int64(-1)

	idStr := chi.URLParam(r, `id`)
	if idStr != `` {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	var dto DTOReoccurringTaskAdd

	if err := readStruct(r.Body, &dto); err != nil {
		panic(err)
	}

	newid, userErrors, internalError := restapi.api.ReOccurringTaskUpdate(context.TODO(),
		serviceitems.ReoccurringTaskUpdate{
			Id:       int(id),
			Title:    dto.Title,
			Duration: time.Second * time.Duration(dto.Seconds),
		})

	if internalError != nil {
		panic(internalError)
		return
	}

	if userErrors != nil {
		b, err := json.Marshal(userErrors)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(b)
		return
	}

	b, err := json.Marshal(&DTONewId{
		Id: newid,
	})
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) reoccurringTaskList(w http.ResponseWriter, r *http.Request) {
	tasks, internalError := restapi.api.ReOccurringTaskList(context.TODO())

	if internalError != nil {
		panic(internalError)
		return
	}

	if len(tasks) == 0 {
		// No data
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert to JSON

	datefmt := `2006-01-02 15:04:05`

	var l []DTOReoccurringTask
	for _, i := range tasks {
		t := DTOReoccurringTask{
			Id:      i.Id,
			AddedAt: i.AddedAt.Format(datefmt),
			Title:   i.Title,
		}

		l = append(l, t)
	}

	b, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) reoccurringTaskGet(w http.ResponseWriter, r *http.Request) {
	panic(`not implemented`)
}
