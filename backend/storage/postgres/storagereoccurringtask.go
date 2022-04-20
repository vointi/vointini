package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointi/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) reoccurringTaskAdd(ctx context.Context, update serviceitems.ReoccurringTaskUpdate) (retid int, internalError error) {
	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
spawn_tasks 
       (title, min_respawn_duration_after_done) VALUES
       ($1,    $2                             )
RETURNING id
`,
		update.Title, update.Duration,
	)

	if internalError != nil {
		return -1, internalError
	}

	return retid, nil

}

func (s StoragePostgreSQL) ReOccurringTaskUpdate(ctx context.Context, update serviceitems.ReoccurringTaskUpdate) (retid int, internalError error) {
	if update.Id == -1 {
		// Add new
		return s.reoccurringTaskAdd(ctx, update)
	}

	cmdtag, internalError := s.db.Exec(ctx, `
UPDATE 
  spawn_tasks 
SET
  title = $2,
  min_respawn_duration_after_done = $3,
WHERE
  id = $1
`,
		update.Id, update.Title, update.Duration,
	)

	if internalError != nil {
		return update.Id, fmt.Errorf(`could not update task #%d %#v - %v`, update.Id, update, internalError)
	}

	if cmdtag.RowsAffected() != 1 {
		return update.Id, fmt.Errorf(`task %d was not updated %#v`, update.Id, update)
	}

	return update.Id, nil
}

func (s StoragePostgreSQL) ReOccurringTaskList(ctx context.Context) (tasks []*serviceitems.ReoccurringTask, internalError error) {
	var sitems []*reoccurringtask

	internalError = pgxscan.Select(ctx, s.db, &sitems,
		`SELECT 
  id
  ,title
  ,added_at 
FROM 
  spawn_tasks
ORDER BY 
  id DESC
`,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: TaskList: %w`, internalError)
	}

	for _, i := range sitems {
		tasks = append(tasks, i.ConvertToAPI())
	}

	sitems = nil // Free memory

	return tasks, nil
}
