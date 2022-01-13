package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"strings"
	"time"
	"unicode"
)

func (r Service) EntryUpdate(ctx context.Context, item serviceitems.EntryUpdate) (userError []UserError, internalError error) {
	item.Activity = strings.TrimSpace(item.Activity)
	item.Description = strings.TrimSpace(item.Description)

	if item.Activity == `` {
		userError = append(userError, UserError{
			Field: `activity`,
			Msg:   `activity cannot be empty`,
		})
	}

	if item.DateTime.Location() != time.UTC {
		userError = append(userError, UserError{
			Field: `datetime`,
			Msg:   `datetime is using non-UTC time zone`,
		})
	}

	if len(item.Tags) == 0 {
		userError = append(userError, UserError{
			Field: `tags`,
			Msg:   `must have at least one tag`,
		})
	}

	if userError != nil {
		return userError, internalError
	}

	internalError = r.storage.EntryUpdate(ctx, item)

	return userError, internalError
}

func (r Service) EntryGet(ctx context.Context, fromtime time.Time, precision time.Duration) (items []*serviceitems.Entry, internalError error) {
	return r.storage.EntryGet(ctx, fromtime, precision)
}

var entryLevelListCache []*serviceitems.EntryLevel

func (r Service) EntryLevelsList(ctx context.Context) (levelList []*serviceitems.EntryLevel, internalError error) {

	if entryLevelListCache == nil {
		entryLevelListCache, internalError = r.storage.EntryLevelsList(ctx)

		if internalError != nil {
			entryLevelListCache = nil
			return nil, internalError
		}
	}

	return entryLevelListCache, nil
}

func (r Service) EntryLevelUpdate(ctx context.Context, update serviceitems.EntryLevelUpdate) (userError []UserError, internalError error) {
	entryLevelListCache = nil

	update.Name = strings.TrimSpace(update.Name)

	update.ShortName = strings.TrimSpace(update.ShortName)
	update.ShortName = strings.ToLower(update.ShortName)

	if update.Name == `` {
		userError = append(userError, UserError{
			Field: "name",
			Msg:   "name cannot be empty",
		})
	}

	if update.ShortName == `` {
		userError = append(userError, UserError{
			Field: "shortname",
			Msg:   "shortname cannot be empty",
		})
	}

	if !asciiLettersOnly(update.ShortName) {
		userError = append(userError, UserError{
			Field: "shortname",
			Msg:   "shortname must be only ASCII letters",
		})
	}

	if userError != nil {
		return userError, nil
	}

	internalError = r.storage.EntryLevelUpdate(ctx, update)
	if internalError != nil {
		return nil, internalError
	}

	return nil, nil
}

func asciiLettersOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}

		if !(r >= 'a' && r <= 'z') {
			return false
		}
	}

	return true
}
