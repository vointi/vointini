package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"strings"
)

var tagCache []*serviceitems.Tag

func (r Service) TagList(ctx context.Context) (tags []*serviceitems.Tag, err error) {
	if tagCache != nil {
		return tagCache, nil
	}

	tags, err = r.storage.TagList(ctx)
	if err != nil {
		tagCache = tags
	}

	return tags, err
}

func (r Service) TagUpdate(ctx context.Context, tag serviceitems.TagUpdate) (newid int, userErrors []UserError, internalError error) {
	tagCache = nil

	tag.Name = strings.TrimSpace(tag.Name)

	tag.ShortName = strings.TrimSpace(tag.ShortName)
	tag.ShortName = strings.ToLower(tag.ShortName)

	if tag.Name == `` {
		userErrors = append(userErrors, UserError{
			Field: "name",
			Msg:   "name is empty",
		})
	}

	if tag.ShortName == `` {
		userErrors = append(userErrors, UserError{
			Field: "shortname",
			Msg:   "shortname is empty",
		})
	}

	if !asciiLettersOnly(tag.ShortName) {
		userErrors = append(userErrors, UserError{
			Field: "shortname",
			Msg:   "shortname must be only letters",
		})
	}

	if userErrors != nil {
		return tag.Id, userErrors, nil
	}

	newid, internalError = r.storage.TagUpdate(ctx, tag)
	return newid, nil, internalError
}

func (r Service) TagGet(ctx context.Context, i int) (*serviceitems.Tag, error) {
	return r.storage.TagGet(ctx, i)
}
