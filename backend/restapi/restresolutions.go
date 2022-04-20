package restapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vointi/vointini/backend/serviceapi"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (restapi restAPI) resolutionsList(w http.ResponseWriter, r *http.Request) {
	entlist, err := restapi.api.ResolutionsEntityList(context.TODO())
	if err != nil {
		panic(err)
		return
	}

	entities := make(map[int]string)
	for _, e := range entlist {
		entities[e.Id] = e.Name
	}

	l, internalError := restapi.api.ResolutionsList(context.TODO())
	if internalError != nil {
		panic(internalError)
		return
	}

	if l == nil {
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert internal format to JSON API format
	var ditems []DTOResolutions

	dateFmt := `2006-01-01`

	for _, i := range l {
		add := DTOResolutions{
			Id:        i.Id,
			EntityId:  i.EntityId,
			AddedAt:   i.AddedAt.String(),
			Name:      i.Name,
			StartDate: i.StartDate.Format(dateFmt),
		}

		if i.SentDate != nil {
			add.SentDate = i.SentDate.Format(dateFmt)
		}

		if i.DecisionDate != nil {
			add.DecisionDate = i.DecisionDate.Format(dateFmt)
		}

		if i.EndDate != nil {
			add.EndDate = i.EndDate.Format(dateFmt)
		}

		// Get file(s) associated with resolution
		files, err := restapi.api.ResolutionsGetFiles(context.TODO(), i.Id)

		if err != nil {
			panic(err)
			return
		}

		for _, f := range files {
			fext := path.Ext(f.Filename)

			add.Files = append(add.Files, DTOResolutionFile{
				Id:      f.Id,
				AddedAt: f.AddedAt.Format(dateFmt),
				Name: fmt.Sprintf(
					`%s %s #%d%s`,
					i.StartDate.Format(dateFmt), entities[i.EntityId], f.Id, fext),
			})
		}

		// Append item
		ditems = append(ditems, add)
	}

	l = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) resolutionsUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	var item DTOResolutionsUpdate
	if err := readStruct(r.Body, &item); err != nil {
		panic(err)
	}

	update := serviceitems.ResolutionsUpdate{
		Id:       id,
		EntityId: item.EntityId,
		Name:     item.Name,
	}

	var userErrors []serviceapi.UserError

	if item.SentDate != `` {
		if tmp, err := time.Parse(`2006-01-02`, item.SentDate); err != nil {
			userErrors = append(userErrors, serviceapi.UserError{
				Field: "sentdate",
				Msg:   restapi.tr.Sprintf(`date.invalid`),
			})
		} else {
			update.SentDate = &tmp
		}
	}

	if item.DecisionDate != `` {
		if tmp, err := time.Parse(`2006-01-02`, item.DecisionDate); err != nil {
			userErrors = append(userErrors, serviceapi.UserError{
				Field: "decisiondate",
				Msg:   restapi.tr.Sprintf(`date.invalid`),
			})
		} else {
			update.DecisionDate = &tmp
		}
	}

	if item.StartDate != `` {
		if tmp, err := time.Parse(`2006-01-02`, item.StartDate); err != nil {
			userErrors = append(userErrors, serviceapi.UserError{
				Field: "startdate",
				Msg:   restapi.tr.Sprintf(`date.invalid`),
			})
		} else {
			update.StartDate = tmp
		}
	} else {
		userErrors = append(userErrors, serviceapi.UserError{
			Field: "startdate",
			Msg:   restapi.tr.Sprintf(`str.empty`),
		})
	}

	if item.EndDate != `` {
		if tmp, err := time.Parse(`2006-01-02`, item.EndDate); err != nil {
			userErrors = append(userErrors, serviceapi.UserError{
				Field: "enddate",
				Msg:   restapi.tr.Sprintf(`date.invalid`),
			})
		} else {
			update.EndDate = &tmp
		}
	}

	if userErrors != nil {
		// got error(s)
		w.WriteHeader(http.StatusBadRequest)
		b, err := json.Marshal(userErrors)
		if err != nil {
			panic(err)
		}

		_, _ = w.Write(b)
		return
	}

	newid, userErrors, internalError := restapi.api.ResolutionsUpdate(context.TODO(), update)

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

func (restapi restAPI) resolutionsGet(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	item, internalError := restapi.api.ResolutionsGet(context.TODO(), id)
	if internalError != nil {
		panic(internalError)
		return
	}

	dateFmt := `2006-01-01`

	dto := DTOResolutions{
		Id:        item.Id,
		EntityId:  item.EntityId,
		AddedAt:   item.AddedAt.String(),
		Name:      item.Name,
		StartDate: item.StartDate.Format(dateFmt),
	}

	if item.DecisionDate != nil {
		dto.DecisionDate = item.DecisionDate.Format(dateFmt)
	}

	if item.SentDate != nil {
		dto.SentDate = item.SentDate.Format(dateFmt)
	}

	if item.EndDate != nil {
		dto.EndDate = item.EndDate.Format(dateFmt)
	}

	b, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) resolutionsEntityList(w http.ResponseWriter, r *http.Request) {
	l, internalError := restapi.api.ResolutionsEntityList(context.TODO())
	if internalError != nil {
		panic(internalError)
		return
	}

	if l == nil {
		// Empty
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert internal format to JSON API format
	var ditems []DTOResolutionEntity

	for _, i := range l {
		ditems = append(ditems, DTOResolutionEntity{
			Id:   i.Id,
			Name: i.Name,
		})
	}

	l = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// resolutionFileUpload uploads a new file to a given resolution.
// Use multipart/form-data MIME type.
func (restapi restAPI) resolutionFileUpload(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`) // Resolution id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	// TODO: check that given resolution ID exists

	var uploadedNames []string

	maxUploadSize := int64(1024 * 1024)

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	if r.MultipartForm == nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(fmt.Errorf(`multipartform is empty`))
		return
	}

	for _, fl := range r.MultipartForm.File {
		// iterate through uploaded file(s)

		if fl == nil || len(fl) == 0 {
			// No file(s)
			continue
		}

		for _, fh := range fl {
			if fh == nil {
				// No file(s)
				continue
			}

			f, err := fh.Open()
			if err != nil {
				panic(err)
				return
			}

			mimetype := strings.ToLower(fh.Header.Get(`Content-Type`))

			// TODO better way to add missing charset?
			switch mimetype {
			case `text/plain`:
				mimetype += `; charset=utf-8`
			}

			// Send stream to be uploaded
			upfname, _, err := restapi.api.ResolutionsUploadFile(context.TODO(), f, id, fh.Filename, mimetype)
			if err != nil {
				panic(err)
				return
			}

			err = f.Close()
			if err != nil {
				panic(err)
				return
			}

			uploadedNames = append(uploadedNames, upfname)
		}

	}

	b, err := json.Marshal(uploadedNames)
	if err != nil {
		panic(err)
		return
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) resolutionFileList(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`) // Resolution id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	_, err = restapi.api.ResolutionsGetFiles(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

}

// resolutionFileGet sends file binary
func (restapi restAPI) resolutionFileGet(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`) // File id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	f, mimeType, err := restapi.api.ResolutionsGetFile(context.TODO(), id)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// No file with this ID
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Internal error
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}
	defer f.Close()

	w.Header().Set(`Content-Type`, mimeType)

	_, _ = io.Copy(w, f)
}
