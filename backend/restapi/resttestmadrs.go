package restapi

import (
	"context"
	"encoding/json"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
	"net/http"
)

func (restapi restAPI) testMADRS(w http.ResponseWriter, r *http.Request) {
	var item DTOTestMADRSAnswers
	if err := readStruct(r.Body, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	userErrors, internalError := restapi.api.TestMADRSAnswer(context.TODO(),
		serviceitems.TestMADRSAnswers{
			Answers: []int{
				item.Answer1,
				item.Answer2,
				item.Answer3,
				item.Answer4,
				item.Answer5,
				item.Answer6,
				item.Answer7,
				item.Answer8,
				item.Answer9,
				item.Answer10,
			},
			Score: 0,
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

func (restapi restAPI) testMADRSResults(w http.ResponseWriter, r *http.Request) {

}
