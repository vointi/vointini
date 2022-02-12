package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"io"
	"path"
	"strings"
	"time"
)

func (r Service) ResolutionsUpdate(ctx context.Context, update serviceitems.ResolutionsUpdate) (newid int, userError []UserError, internalError error) {
	update.Name = strings.TrimSpace(update.Name)

	if update.Name == `` {
		userError = append(userError, UserError{
			Field: `name`,
			Msg:   r.tr.Sprintf(`str.empty`),
		})
	}

	// Validate timezone

	if update.StartDate.Location() != time.UTC {
		userError = append(userError, UserError{
			Field: `startdate`,
			Msg:   r.tr.Sprintf(`datetime.non-UTC`),
		})
	}

	if update.EndDate != nil && update.EndDate.Location() != time.UTC {
		userError = append(userError, UserError{
			Field: `enddate`,
			Msg:   r.tr.Sprintf(`datetime.non-UTC`),
		})
	}

	if update.SentDate != nil && update.SentDate.Location() != time.UTC {
		userError = append(userError, UserError{
			Field: `sentdate`,
			Msg:   r.tr.Sprintf(`datetime.non-UTC`),
		})
	}

	if update.DecisionDate != nil && update.DecisionDate.Location() != time.UTC {
		userError = append(userError, UserError{
			Field: `decisiondate`,
			Msg:   r.tr.Sprintf(`datetime.non-UTC`),
		})
	}

	if userError != nil {
		return update.Id, userError, nil
	}

	newid, internalError = r.storage.ResolutionsUpdate(ctx, update)

	return newid, nil, internalError
}

func (r Service) ResolutionsList(ctx context.Context) (list []*serviceitems.Resolution, err error) {
	return r.storage.ResolutionsList(ctx)
}

func (r Service) ResolutionsEntityList(ctx context.Context) (list []*serviceitems.ResolutionEntity, err error) {
	return r.storage.ResolutionsEntityList(ctx)
}

func (r Service) ResolutionsGet(ctx context.Context, id int) (*serviceitems.Resolution, error) {
	return r.storage.ResolutionsGet(ctx, id)
}

func (r Service) ResolutionsGetFiles(ctx context.Context, resolutionId int) ([]*serviceitems.ResolutionFile, error) {
	return r.storage.ResolutionsGetFiles(ctx, resolutionId)
}

// ResolutionsUploadFile uploads a new file into the filesystem and metadata to storage database
func (r Service) ResolutionsUploadFile(ctx context.Context, src io.ReadCloser, resolutionId int, filename string) (newname string, newid int, err error) {
	fileExtension := `dat` // unknown

	if strings.ContainsRune(filename, '.') {
		fileExtension = strings.TrimLeft(path.Ext(filename), `.`)
	}

	// Add file to a filesystem
	newname, err = r.fileStorage.AddResolutionFile(ctx, resolutionId, fileExtension, src)
	if err != nil {
		return newname, newid, err
	}

	// Add metadata to database
	newid, err = r.storage.ResolutionsFileAdd(ctx, resolutionId, newname)
	if err != nil {
		return newname, newid, err
	}

	return newname, newid, nil
}
