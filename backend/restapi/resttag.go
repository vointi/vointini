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

func (restapi restAPI) tagList(w http.ResponseWriter, r *http.Request) {
	tags, internalError := restapi.api.TagList(context.TODO())

	if internalError != nil {
		panic(internalError)
		return
	}

	if len(tags) == 0 {
		// No data
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert to JSON

	datefmt := `2006-01-02 15:04:05`

	var l []DTOTag
	for _, i := range tags {
		l = append(l, DTOTag{
			Id:        i.Id,
			AddedAt:   i.AddedAt.Format(datefmt),
			Name:      i.Name,
			ShortName: i.ShortName,
		})
	}

	b, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) tagUpdate(w http.ResponseWriter, r *http.Request) {
	var err error

	id, err := getIntParam(r, `id`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	var item DTOTag
	if err := readStruct(r.Body, &item); err != nil {
		panic(err)
	}

	newid, userErrors, internalError := restapi.api.TagUpdate(context.TODO(),
		serviceitems.TagUpdate{
			Id:        id,
			Name:      item.Name,
			ShortName: item.ShortName,
		})

	if internalError != nil {
		panic(internalError)
		return
	}

	if userErrors != nil {
		w.WriteHeader(http.StatusBadRequest)
		b, err := json.Marshal(userErrors)
		if err != nil {
			panic(err)
		}

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

func (restapi restAPI) tagGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, `id`), 10, 64)
	if err != nil {
		panic(err)
	}

	tag, internalError := restapi.api.TagGet(context.TODO(), int(id))

	if internalError != nil {
		panic(internalError)
		return
	}

	if tag == nil {
		// No data
		_, _ = io.WriteString(w, `{}`)
		return
	}

	// Convert to JSON

	t := DTOTag{
		Id:        tag.Id,
		Name:      tag.Name,
		ShortName: tag.ShortName,
	}

	b, err := json.Marshal(&t)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) convertTagsFromInternal(from []int) (to []string) {
	list, err := restapi.api.TagList(context.TODO())
	if err != nil {
		return nil
	}

	for _, i := range list {
		for _, id := range from {
			if id == i.Id {
				to = append(to, i.ShortName)
			}
		}
	}

	return to
}

func (restapi restAPI) convertTagsFromHuman(from []string) (to []int) {
	list, err := restapi.api.TagList(context.TODO())
	if err != nil {
		return nil
	}

	for _, i := range list {
		for _, id := range from {
			if id == i.ShortName {
				to = append(to, i.Id)
			}
		}
	}

	return to

}
