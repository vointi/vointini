package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (r Service) TaskUpdate(ctx context.Context, task serviceitems.TaskUpdate) (id int, userErrors []UserError, internalError error) {
	if task.Title == `` {
		userErrors = append(userErrors, UserError{
			Field: "title",
			Msg:   "title is required",
		})
	}

	if userErrors != nil {
		return -1, userErrors, nil
	}

	newid, internalError := r.storage.TaskUpdate(ctx, task)
	if internalError != nil {
		return id, nil, internalError
	}

	return newid, nil, nil
}

func (r Service) TaskList(ctx context.Context, option serviceitems.TaskSearchOption) (tasks []*serviceitems.Task, internalError error) {
	return r.storage.TaskList(ctx, option)
}

func (r Service) TaskGet(ctx context.Context, id int) (task *serviceitems.Task, internalError error) {
	return r.storage.TaskGet(ctx, id)
}
