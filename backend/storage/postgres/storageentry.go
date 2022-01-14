package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"time"
)

func (s StoragePostgreSQL) getEntryId(ctx context.Context, fromtime time.Time) (id int, internalError error) {
	var sitem []*entryId
	internalError = pgxscan.Select(ctx, s.db, &sitem,
		`SELECT 
  id
FROM 
  entries
WHERE 
  date_trunc('minute', fromtime) = date_trunc('minute', $1::timestamp)
LIMIT 1
`,
		fromtime.Truncate(time.Minute*1),
	)

	if internalError != nil {
		return -1, fmt.Errorf(`pg: getEntryId: %w`, internalError)
	}

	if len(sitem) == 1 {
		return sitem[0].Id, nil
	}

	return -1, fmt.Errorf(`pg: getEntryId: not found %v`, fromtime)
}

func (s StoragePostgreSQL) getLevelsOfEntry(ctx context.Context, entryId int) (levels map[int]int, internalError error) {
	levels = make(map[int]int)

	var sitems []*entrysLevel

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  levelid, level 
FROM 
  entry_levels
WHERE 
  entryid=$1
`,
		entryId,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: getLevelsOfEntry: %w`, internalError)
	}

	for _, v := range sitems {
		levels[v.Id] = v.Level
	}

	sitems = nil // Free memory

	return levels, nil
}

func (s StoragePostgreSQL) EntryGet(ctx context.Context, fromtime time.Time, precision time.Duration) (entries []*serviceitems.Entry, internalError error) {
	var sitems []*entry

	truncate := `day`

	if precision.Minutes() == 1.0 {
		truncate = `minute`
	}

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id 
  ,fromtime 
  ,activity_name
  ,level_achievement
  ,descr
FROM 
  entries
WHERE 
  deleted_at IS NULL AND
  date_trunc($1, fromtime) = date_trunc($1, $2::timestamp)
ORDER BY 
  fromtime DESC
`,
		truncate, fromtime.Truncate(precision),
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: EntryGet: %w`, internalError)
	}

	for _, i := range sitems {
		e := i.ConvertToAPI()

		e.Levels, internalError = s.getLevelsOfEntry(ctx, i.Id)
		if internalError != nil {
			return nil, fmt.Errorf(`pg: EntryGet: %w`, internalError)
		}

		e.Tags, internalError = s.getTagsOfEntry(ctx, i.Id)
		if internalError != nil {
			return nil, fmt.Errorf(`pg: EntryGet: %w`, internalError)
		}

		entries = append(entries, e)
	}

	sitems = nil // Free memory

	return entries, nil
}

func (s StoragePostgreSQL) EntryUpdate(ctx context.Context, item serviceitems.EntryUpdate) (internalError error) {

	internalError = s.db.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var eId int

		if len(item.Levels) == 0 {
			return fmt.Errorf(`no levels given`)
		}

		internalError = tx.QueryRow(ctx,
			`INSERT INTO 
entries 
	(fromtime, activity_name, level_achievement) VALUES 
	($1,       $2,            $3               )
ON CONFLICT (fromtime) DO UPDATE SET
  modified_at = now()
  ,activity_name = $2
  ,level_achievement = $3
RETURNING id
`,
			item.DateTime, item.Activity, item.LevelAchievement,
		).Scan(&eId)

		if internalError != nil {
			return internalError
		}

		if eId == 0 {
			return fmt.Errorf(`id is zero`)
		}

		for lvl, setLevel := range item.Levels {
			ct, internalError := tx.Exec(ctx,
				`INSERT INTO 
entry_levels 
	(entryid, levelid, level) VALUES 
	($1,      $2,      $3   )
ON CONFLICT (entryid, levelid) DO UPDATE SET
  level = $3
`,
				eId, lvl, setLevel,
			)

			if internalError != nil {
				return internalError
			}

			if ct.RowsAffected() != 1 {
				return tx.Rollback(ctx)
			}

		}

		// Remove possible old tag(s) first
		_, internalError := tx.Exec(ctx,
			`DELETE
FROM 
  tags_for_entry 
WHERE 
  entryid = $1
`,
			eId,
		)
		if internalError != nil {
			return internalError
		}

		// Add tags
		for _, tagId := range item.Tags {
			ct, internalError := tx.Exec(ctx,
				`INSERT INTO 
tags_for_entry 
	(entryid, tagid) VALUES 
	($1,      $2   )
`,
				eId, tagId,
			)

			if internalError != nil {
				return internalError
			}

			if ct.RowsAffected() != 1 {
				return tx.Rollback(ctx)
			}

		}

		return nil
	})

	if internalError != nil {
		return internalError
	}

	return nil
}

func (s StoragePostgreSQL) EntryLevelsList(ctx context.Context) (entries []*serviceitems.EntryLevel, internalError error) {
	var sitems []*entryLevel

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,name
  ,default_show
  ,worst_descr
  ,shortname
  ,get_previous 
  ,added_at
FROM 
  levels
ORDER BY 
  name ASC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: EntryLevelsList: %w`, internalError)
	}

	for _, i := range sitems {
		entries = append(entries, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return entries, nil
}

func (s StoragePostgreSQL) EntryLevelUpdate(ctx context.Context, update serviceitems.EntryLevelUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (s StoragePostgreSQL) getTagsOfEntry(ctx context.Context, entryId int) (tagIds []int, internalError error) {
	var sitems []*entrysTags

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  tagid
FROM 
  tags_for_entry
WHERE
  entryid = $1
`,
		entryId,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: getTagsOfEntry: %w`, internalError)
	}

	for _, i := range sitems {
		tagIds = append(tagIds, i.TagId)
	}

	sitems = nil // Free memory

	return tagIds, nil
}
