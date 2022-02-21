package serviceapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"io"
	"os"
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
func (r Service) ResolutionsUploadFile(ctx context.Context, src io.ReadCloser, resolutionId int, filename string, contentType string) (newname string, newid int, err error) {
	contentType = strings.ToLower(contentType)
	contentType = strings.TrimSpace(contentType)

	filename = strings.TrimSpace(filename)

	if filename == `` {
		return ``, -1, fmt.Errorf(r.tr.Sprintf(`empty.filename`))
	}

	if contentType == `` {
		return ``, -1, fmt.Errorf(r.tr.Sprintf(`empty.content-type`))
	}

	fileExtension := `dat` // unknown file extension

	if strings.ContainsRune(filename, '.') {
		fileExtension = strings.TrimLeft(path.Ext(filename), `.`)
	}

	// Add file to a filesystem
	newname, err = r.fileStorage.AddResolutionFile(ctx, resolutionId, fileExtension, src)
	if err != nil {
		return newname, newid, err
	}

	// Add metadata to database
	newid, err = r.storage.ResolutionsFileAdd(ctx, resolutionId, newname, contentType)
	if err != nil {
		return newname, newid, err
	}

	return newname, newid, nil
}

// ResolutionsGetFile opens a given file from the filesystem
func (r Service) ResolutionsGetFile(ctx context.Context, id int) (src io.ReadCloser, mime string, err error) {
	// Check database that the file reference exists
	file, err := r.storage.ResolutionsGetFile(ctx, id)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, mime, err
		} else {
			return nil, mime, fmt.Errorf(`db err: %w`, err)
		}
	}

	// Open file in filesystem
	src, err = r.fileStorage.GetResolutionFile(ctx, file.ResolutionId, file.Filename)
	if err != nil {
		return nil, mime, fmt.Errorf(`could not open file %d for resolution %d`, id, file.ResolutionId)
	}

	return src, file.ContentType, nil
}
