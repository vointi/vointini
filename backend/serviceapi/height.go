package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (r Service) HeightUpdate(ctx context.Context, add serviceitems.HeightAdd) (userError []UserError, internalError error) {
	if add.Height == 0.0 {
		userError = append(userError, UserError{
			Field: "height",
			Msg:   "height must be non-zero",
		})
	} else if add.Height < 0.0 {
		userError = append(userError, UserError{
			Field: "height",
			Msg:   "height must be non-negative",
		})
	}

	if userError != nil {
		return userError, nil
	}

	internalError = r.storage.HeightUpdate(ctx, add)

	if internalError != nil {
		return nil, internalError
	}

	return nil, nil
}

func (r Service) HeightList(ctx context.Context) (l []*serviceitems.Height, internalError error) {
	return r.storage.HeightList(ctx)
}
