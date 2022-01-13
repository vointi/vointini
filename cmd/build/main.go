package main

import (
	"fmt"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"github.com/vointini/vointini/backend/restapi"
	"os"
	"path"
)

func main() {
	converter := typescriptify.New()
	converter.CreateFromMethod = false
	converter.DontExport = false
	converter.BackupDir = ``
	converter.CreateInterface = true

	// Generic responses
	converter.Add(restapi.DTOOK{})
	converter.Add(restapi.DTONewId{})

	// Entry
	converter.Add(restapi.DTOEntry{})
	converter.Add(restapi.DTOEntryAdd{})
	converter.Add(restapi.DTOEntryLevel{})
	converter.Add(restapi.DTOEntryLevelUpdate{})

	// Tag
	converter.Add(restapi.DTOTag{})

	// Task
	converter.Add(restapi.DTOTask{})

	// Timer
	converter.Add(restapi.DTOTimer{})
	converter.Add(restapi.DTOTimerAdd{})

	// Re-occurring task(s)
	converter.Add(restapi.DTOReoccurringTaskAdd{})
	converter.Add(restapi.DTOReoccurringTask{})

	// Weight
	converter.Add(restapi.DTOWeight{})
	converter.Add(restapi.DTOWeightAdd{})

	// Height
	converter.Add(restapi.DTOHeight{})
	converter.Add(restapi.DTOHeightAdd{})

	// Tests

	// MADRS
	converter.Add(restapi.DTOTestMADRSAnswers{})

	err := converter.ConvertToFile(path.Join(`frontend`, `templates`, `src`, `dto.ts`))
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}
}
