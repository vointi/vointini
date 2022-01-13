package restapi

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"io"
	"net/http"
	"strconv"
)

// Task related

func (restapi restAPI) taskUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	id := int64(-1)

	idStr := chi.URLParam(r, `id`)
	if idStr != `` {
		id, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	var item DTOTask
	if err := readStruct(r.Body, &item); err != nil {
		panic(err)
	}

	newid, userErrors, internalError := restapi.api.TaskUpdate(context.TODO(),
		serviceitems.TaskUpdate{
			Id:          int(id),
			Title:       item.Title,
			Description: item.Description,
		})

	if internalError != nil {
		panic(internalError)
		return
	}

	if len(userErrors) > 0 {

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

func (restapi restAPI) taskList(w http.ResponseWriter, r *http.Request) {
	tasks, internalError := restapi.api.TaskList(context.TODO(), serviceitems.OngoingTasks)

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

	var l []DTOTask
	for _, i := range tasks {
		t := DTOTask{
			Id:                         i.Id,
			AddedAt:                    i.AddedAt.Format(datefmt),
			Title:                      i.Title,
			Description:                i.Description,
			ReoccurringTaskReferenceId: i.ReoccurringTaskReferenceId,
		}

		if i.CompletedAt != nil {
			v := i.CompletedAt.Format(datefmt)
			t.CompletedAt = &v
		}

		l = append(l, t)
	}

	b, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) taskGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, `id`), 10, 64)
	if err != nil {
		panic(err)
	}

	task, internalError := restapi.api.TaskGet(context.TODO(), int(id))

	if internalError != nil {
		panic(internalError)
		return
	}

	if task == nil {
		// No data
		_, _ = io.WriteString(w, ``)
		return
	}

	// Convert to JSON

	t := DTOTask{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
	}

	b, err := json.Marshal(&t)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}
