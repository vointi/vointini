package serviceapi

import (
	"context"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
)

func (r Service) WeightUpdate(ctx context.Context, w serviceitems.WeightAdd) (userError []UserError, internalError error) {
	if w.Weight == 0.0 {
		userError = append(userError, UserError{
			Field: "weight",
			Msg:   "weight must be non-zero",
		})
	} else if w.Weight < 0.0 {
		userError = append(userError, UserError{
			Field: "weight",
			Msg:   "weight must be non-negative",
		})
	}

	if userError != nil {
		return userError, nil
	}

	internalError = r.storage.WeightUpdate(ctx, w)

	if internalError != nil {
		return nil, internalError
	}

	return nil, nil
}

func (r Service) WeightList(ctx context.Context) (l []*serviceitems.Weight, internalError error) {
	return r.storage.WeightList(ctx)
}
