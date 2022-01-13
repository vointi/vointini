package restapi

import (
	"context"
	"encoding/json"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"io"
	"net/http"
)

func (restapi restAPI) heightList(w http.ResponseWriter, r *http.Request) {
	l, internalError := restapi.api.HeightList(context.TODO())
	if internalError != nil {
		panic(internalError)
		return
	}

	if l == nil {
		// no weights added yet
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert internal format to JSON API format
	var ditems []DTOHeight

	for _, i := range l {
		ditems = append(ditems, DTOHeight{
			Height: i.Height,
			Added:  i.AddedAt.Format(`2006-01-02T15:04`),
			Id:     i.Id,
		})
	}

	l = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) heightUpdate(w http.ResponseWriter, r *http.Request) {
	var item DTOHeightAdd
	if err := readStruct(r.Body, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	userErrors, internalError := restapi.api.HeightUpdate(context.TODO(),
		serviceitems.HeightAdd{
			Height: item.Height,
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

	b, err := json.Marshal(&DTOOK{
		Msg: `ok`,
	})

	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}
