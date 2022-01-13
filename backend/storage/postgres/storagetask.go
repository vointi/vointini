package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

func (s StoragePostgreSQL) taskAdd(ctx context.Context, task serviceitems.TaskUpdate) (retid int, internalError error) {
	internalError = pgxscan.Get(ctx, s.db, &retid,
		`INSERT INTO 
tasks 
  (title, descr, priority, refid) VALUES
  ($1,    $2,    $3,       $4   )
RETURNING id
`,
		task.Title, task.Description, task.Priority, task.ReoccurringTaskReferenceId,
	)

	if internalError != nil {
		return -1, internalError
	}

	return retid, nil
}

func (s StoragePostgreSQL) TaskUpdate(ctx context.Context, task serviceitems.TaskUpdate) (retid int, internalError error) {

	if task.Id == -1 {
		// Add new
		return s.taskAdd(ctx, task)
	}

	cmdtag, internalError := s.db.Exec(ctx, `
UPDATE 
  tasks 
SET
  title = $2,
  descr = $3,
  priority = $4
WHERE
  id = $1
`,
		task.Id, task.Title, task.Description, task.Priority,
	)

	if internalError != nil {
		return task.Id, fmt.Errorf(`could not update task #%d %#v - %v`, task.Id, task, internalError)
	}

	if cmdtag.RowsAffected() != 1 {
		return task.Id, fmt.Errorf(`task %d was not updated %#v`, task.Id, task)
	}

	return task.Id, nil
}

func (s StoragePostgreSQL) TaskList(ctx context.Context, option serviceitems.TaskSearchOption) (tasks []*serviceitems.Task, internalError error) {
	var sitems []*task

	sql := `SELECT 
  id
  ,title
  ,descr
  ,added_at
  ,completed_at
  ,refid 
FROM 
  tasks
WHERE
`

	switch option {
	case serviceitems.OngoingTasks:
		sql += `completed_at IS NULL`
	case serviceitems.CompletedTasks:
		sql += `completed_at IS NOT NULL`
	}

	sql += `
ORDER BY 
`

	switch option {
	case serviceitems.OngoingTasks:
		sql += `added_at ASC`
	case serviceitems.CompletedTasks:
		sql += `completed_at DESC`
	}

	internalError = pgxscan.Select(ctx, s.db, &sitems, sql)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: TaskList: %w`, internalError)
	}

	for _, i := range sitems {
		tasks = append(tasks, &serviceitems.Task{
			Id:                         i.Id,
			AddedAt:                    i.AddedAt,
			CompletedAt:                i.CompletedAt,
			Priority:                   i.Priority,
			Title:                      i.Title,
			Description:                i.Description,
			ReoccurringTaskReferenceId: i.ReoccurringTaskReferenceId,
		})
	}

	sitems = nil // Free memory

	return tasks, nil
}

func (s StoragePostgreSQL) TaskGet(ctx context.Context, id int) (item *serviceitems.Task, internalError error) {
	var sitem []*task

	internalError = pgxscan.Select(ctx, s.db, &sitem,
		`SELECT 
  id
  ,title
  ,descr
  ,added_at
  ,completed_at
  ,refid 
FROM 
  tasks
WHERE
  id = $1
LIMIT 1
`,
		id,
	)

	if internalError != nil {
		return nil, fmt.Errorf(`pg: TaskList: %w`, internalError)
	}

	if len(sitem) == 1 {
		item = &serviceitems.Task{
			Id:                         sitem[0].Id,
			AddedAt:                    sitem[0].AddedAt,
			Priority:                   sitem[0].Priority,
			Title:                      sitem[0].Title,
			Description:                sitem[0].Description,
			CompletedAt:                sitem[0].CompletedAt,
			ReoccurringTaskReferenceId: sitem[0].ReoccurringTaskReferenceId,
		}
	}

	sitem = nil // Free memory

	return item, nil
}
