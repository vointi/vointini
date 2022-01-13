package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) TestMADRSAnswer(ctx context.Context, answer serviceitems.TestMADRSAnswers) (internalError error) {
	var retid int

	a := struct {
		A1    int
		A2    int
		A3    int
		A4    int
		A5    int
		A6    int
		A7    int
		A8    int
		A9    int
		A10   int
		Score int
	}{
		A1:    answer.Answers[0],
		A2:    answer.Answers[1],
		A3:    answer.Answers[2],
		A4:    answer.Answers[3],
		A5:    answer.Answers[4],
		A6:    answer.Answers[5],
		A7:    answer.Answers[6],
		A8:    answer.Answers[7],
		A9:    answer.Answers[8],
		A10:   answer.Answers[9],
		Score: answer.Score,
	}

	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
test_madrs 
  (a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, score) VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id
`,
		a.A1, a.A2, a.A3, a.A4, a.A5, a.A6, a.A7, a.A8, a.A9, a.A10, a.Score,
	)

	if internalError != nil {
		return fmt.Errorf(`pg: TestMADRSAnswer: %w`, internalError)
	}

	return nil
}
