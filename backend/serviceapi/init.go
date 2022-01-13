package serviceapi

import (
	"context"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
)

// Initialize adds initial data to Storage
func (r Service) Initialize(ctx context.Context) (internalError error) {
	internalError = r.initLevels(ctx)
	if internalError != nil {
		return internalError
	}

	return nil
}

// initLevels adds initial levels to Storage
func (r Service) initLevels(ctx context.Context) (internalError error) {
	type kvS struct {
		Name string
		Key  string
	}

	levels, err := r.EntryLevelsList(ctx)
	if err != nil {
		return err
	}

	if len(levels) > 0 {
		// Already added
		return nil
	}

	// Initial data to be added
	initLevels := []kvS{
		{
			Name: r.tr.Sprintf(`depression`),
			Key:  `depression`,
		},
		{
			Name: r.tr.Sprintf(`anxiety`),
			Key:  `anxiety`,
		},
		{
			Name: r.tr.Sprintf(`stress`),
			Key:  `stress`,
		},
	}

	for _, lvl := range initLevels {
		// Add to storage
		_, err = r.EntryLevelUpdate(ctx, serviceitems.EntryLevelUpdate{
			Id:        -1,
			Name:      lvl.Name,
			ShortName: lvl.Key,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
