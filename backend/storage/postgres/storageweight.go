package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) WeightUpdate(ctx context.Context, add serviceitems.WeightAdd) (internalError error) {
	var retid int

	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
weight 
  (value) VALUES
  ($1    )
RETURNING id
`,
		add.Weight,
	)

	if internalError != nil {
		return fmt.Errorf(`pg: WeightUpdate: %w`, internalError)

	}

	return nil
}

func (s StoragePostgreSQL) WeightList(ctx context.Context) (items []*serviceitems.Weight, internalError error) {
	var sitems []*weight

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,added_at
  ,value
FROM 
  weight
ORDER BY
  added_at DESC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: WeightList: %w`, internalError)
	}

	for _, i := range sitems {
		items = append(items, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return items, nil
}
