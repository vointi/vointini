package storageimpl

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"time"
)

// Storage interface tells what backend storage (for example: SQL database, filesystem, memory, etc.) must implement
type Storage interface {

	// Entries

	EntryUpdate(context.Context, serviceitems.EntryUpdate) error
	EntryGet(context.Context, time.Time, time.Duration) ([]*serviceitems.Entry, error)
	EntryLevelsList(context.Context) ([]*serviceitems.EntryLevel, error)
	EntryLevelUpdate(context.Context, serviceitems.EntryLevelUpdate) error

	// Tags

	TagList(ctx context.Context) ([]*serviceitems.Tag, error)
	TagUpdate(ctx context.Context, tag serviceitems.TagUpdate) (int, error)
	TagGet(ctx context.Context, i int) (*serviceitems.Tag, error)

	// Tasks

	TaskUpdate(context.Context, serviceitems.TaskUpdate) (int, error)
	TaskList(context.Context, serviceitems.TaskSearchOption) ([]*serviceitems.Task, error)
	TaskGet(context.Context, int) (*serviceitems.Task, error)

	// Re-occurring tasks

	ReOccurringTaskUpdate(context.Context, serviceitems.ReoccurringTaskUpdate) (int, error)
	ReOccurringTaskList(ctx context.Context) ([]*serviceitems.ReoccurringTask, error)

	// Weight

	WeightUpdate(context.Context, serviceitems.WeightAdd) error
	WeightList(context.Context) ([]*serviceitems.Weight, error)

	// Height

	HeightUpdate(context.Context, serviceitems.HeightAdd) error
	HeightList(context.Context) ([]*serviceitems.Height, error)

	// Tests

	// TestMADRSAnswer answers for test: MADRS
	TestMADRSAnswer(context.Context, serviceitems.TestMADRSAnswers) error
}
