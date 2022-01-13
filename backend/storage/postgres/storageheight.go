package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) HeightUpdate(ctx context.Context, add serviceitems.HeightAdd) (internalError error) {
	var retid int

	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
height 
  (value) VALUES
  ($1   )
RETURNING id
`,
		add.Height,
	)

	if internalError != nil {
		return fmt.Errorf(`pg: WeightUpdate: %w`, internalError)

	}

	return nil
}

func (s StoragePostgreSQL) HeightList(ctx context.Context) (items []*serviceitems.Height, internalError error) {
	var sitems []*height

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,added_at
  ,value
FROM 
  height
ORDER BY
  added_at DESC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: WeightList: %w`, internalError)
	}

	for _, i := range sitems {
		items = append(items, &serviceitems.Height{
			Id:      i.Id,
			AddedAt: i.AddedAt,
			Height:  i.Height,
		})
	}

	sitems = nil // Free memory

	return items, nil
}
