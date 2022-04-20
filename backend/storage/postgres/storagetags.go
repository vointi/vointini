package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) TagList(ctx context.Context) (tags []*serviceitems.Tag, internalError error) {
	var sitems []*tag

	sql := `SELECT 
  id
  ,added_at
  ,name
  ,shortname
FROM 
  entry_tags
ORDER BY
  name
`

	internalError = pgxscan.Select(ctx, s.db, &sitems, sql)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: TaskList: %w`, internalError)
	}

	for _, i := range sitems {
		tags = append(tags, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return tags, nil
}

func (s StoragePostgreSQL) tagAdd(ctx context.Context, tag serviceitems.TagUpdate) (retid int, internalError error) {
	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
entry_tags 
  (name, shortname) VALUES
  ($1,   $2       )
RETURNING id
`,
		tag.Name, tag.ShortName,
	)

	if internalError != nil {
		return -1, internalError
	}

	return retid, nil
}

func (s StoragePostgreSQL) TagUpdate(ctx context.Context, tag serviceitems.TagUpdate) (int, error) {
	if tag.Id == -1 {
		return s.tagAdd(ctx, tag)
	}

	cmdtag, internalError := s.db.Exec(ctx, `
UPDATE 
  entry_tags 
SET
  name = $2
  ,shortname = $3
WHERE
  id = $1
`,
		tag.Id, tag.Name, tag.ShortName,
	)

	if internalError != nil {
		return tag.Id, fmt.Errorf(`could not update tag #%d %#v - %v`, tag.Id, tag, internalError)
	}

	if cmdtag.RowsAffected() != 1 {
		return tag.Id, fmt.Errorf(`tag %d was not updated %#v`, tag.Id, tag)
	}

	return tag.Id, nil
}

func (s StoragePostgreSQL) TagGet(ctx context.Context, i int) (item *serviceitems.Tag, internalError error) {
	var sitems []*tag

	sql := `SELECT 
  id
  ,added_at
  ,name
  ,shortname
FROM 
  entry_tags
WHERE
  id = $1
LIMIT 1
`

	internalError = pgxscan.Select(ctx, s.db, &sitems, sql, i)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: TagGet: %w`, internalError)
	}

	if len(sitems) == 1 {
		item = sitems[0].ConvertToAPI()
	}

	sitems = nil // Free memory

	return item, nil
}
