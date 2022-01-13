package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (r Service) TestMADRSAnswer(ctx context.Context, answers serviceitems.TestMADRSAnswers) (userError []UserError, internalError error) {

	if len(answers.Answers) != 10 {
		userError = append(userError, UserError{
			Field: "answers",
			Msg:   "answers length must be 10",
		})
	}

	for _, a := range answers.Answers {
		switch a {
		case 0:
			answers.Score += 0
		case 1:
			answers.Score += 2
		case 2:
			answers.Score += 4
		case 3:
			answers.Score += 6
		}
	}

	if userError != nil {
		return userError, nil
	}

	// Save
	internalError = r.storage.TestMADRSAnswer(ctx, answers)

	return nil, internalError
}
