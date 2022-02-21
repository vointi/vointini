package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"os"
)

func (s StoragePostgreSQL) resolutionsAdd(ctx context.Context, add serviceitems.ResolutionsUpdate) (retid int, internalError error) {
	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
resolutions
  (name, entityid, sentdate, decisiondate, startdate, enddate) VALUES
  ($1,   $2,       $3,       $4,           $5,        $6     )
RETURNING id
`,
		add.Name, add.EntityId, add.SentDate, add.DecisionDate, add.StartDate, add.EndDate,
	)

	if internalError != nil {
		return -1, internalError
	}

	return retid, nil
}

func (s StoragePostgreSQL) ResolutionsUpdate(ctx context.Context, update serviceitems.ResolutionsUpdate) (int, error) {
	if update.Id == -1 {
		return s.resolutionsAdd(ctx, update)
	}

	cmdtag, internalError := s.db.Exec(ctx, `
UPDATE
  resolutions
SET
  name = $2
  ,entityid = $3
  ,sentdate = $4
  ,decisiondate = $5
  ,startdate = $6
  ,enddate = $7
WHERE
  id = $1
`,
		update.Id, update.Name, update.EntityId, update.SentDate, update.DecisionDate, update.StartDate, update.EndDate,
	)

	if internalError != nil {
		return update.Id, fmt.Errorf(`could not update resolution #%d %#v - %v`, update.Id, update, internalError)
	}

	if cmdtag.RowsAffected() != 1 {
		return update.Id, fmt.Errorf(`resolution %d was not updated %#v`, update.Id, update)
	}

	return update.Id, nil

}

func (s StoragePostgreSQL) ResolutionsList(ctx context.Context) (items []*serviceitems.Resolution, internalError error) {
	var sitems []*resolution

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,entityid
  ,added_at
  ,name
  ,decisiondate
  ,sentdate
  ,startdate
  ,enddate
FROM 
  resolutions
ORDER BY
  enddate DESC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: ResolutionsList: %w`, internalError)
	}

	for _, i := range sitems {
		items = append(items, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return items, nil
}

func (s StoragePostgreSQL) ResolutionsEntityList(ctx context.Context) (items []*serviceitems.ResolutionEntity, internalError error) {
	var sitems []*resolutionEntity

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,name
FROM 
  resolution_entity
ORDER BY
  name ASC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: ResolutionsEntityList: %w`, internalError)
	}

	for _, i := range sitems {
		items = append(items, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return items, nil
}

func (s StoragePostgreSQL) ResolutionsGet(ctx context.Context, id int) (item *serviceitems.Resolution, internalError error) {
	var sitems []*resolution

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,entityid
  ,added_at
  ,name
  ,decisiondate
  ,sentdate
  ,startdate
  ,enddate
FROM 
  resolutions
WHERE 
  id = $1
`,
		id,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: ResolutionsGet: %w`, internalError)
	}

	if len(sitems) == 1 {
		item = sitems[0].ConvertToAPI()
	}

	sitems = nil // Free memory

	return item, nil
}

func (s StoragePostgreSQL) ResolutionsGetFiles(ctx context.Context, resolutionId int) (items []*serviceitems.ResolutionFile, internalError error) {
	var sitems []*resolutionFile

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,added_at
  ,resolutionid
  ,filename
FROM 
  resolution_files
WHERE 
  resolutionid = $1
`,
		resolutionId,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: ResolutionsGet: %w`, internalError)
	}

	for _, i := range sitems {
		items = append(items, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return items, nil
}

func (s StoragePostgreSQL) ResolutionsFileAdd(ctx context.Context, resolutionid int, filename string, contentType string) (retid int, internalError error) {
	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
resolution_files
  (resolutionid, filename, ctype) VALUES
  ($1,           $2,       $3   )
ON CONFLICT (filename) DO UPDATE SET
  resolutionid = $1
  ,filename = $2
  ,ctype = $3
RETURNING id
`,
		resolutionid, filename, contentType,
	)

	if internalError != nil {
		return -1, internalError
	}

	return retid, nil
}

func (s StoragePostgreSQL) ResolutionsGetFile(ctx context.Context, id int) (item *serviceitems.ResolutionFile, internalError error) {
	var sitems []*resolutionFile

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,added_at
  ,resolutionid
  ,filename
  ,ctype
FROM 
  resolution_files
WHERE 
  id = $1
LIMIT 1
`,
		id,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: ResolutionsGetFile: %w`, internalError)
	}

	if sitems == nil || len(sitems) != 1 {
		return nil, os.ErrNotExist
	}

	return sitems[0].ConvertToAPI(), nil
}
