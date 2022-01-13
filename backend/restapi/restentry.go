package restapi

import (
	"context"
	"encoding/json"
	"github.com/vointini/vointini/backend/serviceapi/serviceitems"
	"io"
	"net/http"
	"time"
)

// Entry related

func (restapi restAPI) levelToId() (levelNameToId map[string]int, levelIdToName map[int]string, err error) {
	levelNameToId = make(map[string]int)
	levelIdToName = make(map[int]string)

	lvls, err := restapi.api.EntryLevelsList(context.TODO())

	if err != nil {
		return levelNameToId, levelIdToName, err
	}

	if len(lvls) == 0 {
		return levelNameToId, levelIdToName, nil
	}

	for _, i := range lvls {
		levelNameToId[i.ShortName] = i.Id
		levelIdToName[i.Id] = i.ShortName
	}

	return levelNameToId, levelIdToName, nil
}

// convertLevelsFromInternal converts internal level keys as integers -> level keys as strings
func (restapi *restAPI) convertLevelsFromInternal(levels map[int]int) (m map[string]int) {
	m = make(map[string]int)

	_, levelToId, err := restapi.levelToId()
	if err != nil {
		panic(err)
	}

	for _, v := range levelToId {
		m[v] = 0
	}

	for k, v := range levels {
		m[levelToId[k]] = v
	}

	return m
}

// convertLevelsFromHuman converts human level keys as strings -> level keys as integers
func (restapi *restAPI) convertLevelsFromHuman(levels map[string]int) (m map[int]int) {
	m = make(map[int]int)

	levelNameToId, _, internalError := restapi.levelToId()
	if internalError != nil {
		panic(internalError)
		return
	}

	for k, v := range levels {
		m[levelNameToId[k]] = v
	}

	return m
}

// entryUpdate updates single entry
func (restapi restAPI) entryUpdate(w http.ResponseWriter, r *http.Request) {
	year, err := getIntParam(r, `year`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	month, err := getIntParam(r, `month`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	day, err := getIntParam(r, `day`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	hour, err := getIntParam(r, `hour`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	minute, err := getIntParam(r, `minute`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	var item DTOEntryAdd
	if err := readStruct(r.Body, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	userErrors, internalError := restapi.api.EntryUpdate(context.TODO(),
		serviceitems.EntryUpdate{
			Activity:         item.Activity,
			Description:      item.Description,
			DateTime:         time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
			LevelAchievement: item.LevelAchievement,
			Levels:           restapi.convertLevelsFromHuman(item.Levels),
			Tags:             restapi.convertTagsFromHuman(item.Tags),
		})

	if internalError != nil {
		panic(internalError)
		return
	}

	if userErrors != nil {
		b, err := json.Marshal(userErrors)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(b)
		return
	}

	b, err := json.Marshal(&DTOOK{
		Msg: `ok`,
	})

	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// entriesMinute fetches entries with minute accuracy (yyyy-mm-dd hh:mm)
func (restapi restAPI) entriesMinute(w http.ResponseWriter, r *http.Request) {
	year, err := getIntParam(r, `year`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	month, err := getIntParam(r, `month`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	day, err := getIntParam(r, `day`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	hour, err := getIntParam(r, `hour`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	minute, err := getIntParam(r, `minute`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	fromtime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	precision := time.Minute * 1

	eItems, internalError := restapi.api.EntryGet(context.TODO(), fromtime, precision)

	if internalError != nil {
		panic(internalError)
	}

	if len(eItems) == 0 {
		// No data
		_, _ = io.WriteString(w, `{}`)
		return
	}

	// Convert internal format to JSON API format
	ditems := make(map[string]DTOEntry)

	for _, i := range eItems {
		ditems[i.DateTime.Format(`15:04`)] = DTOEntry{
			Id:               i.Id,
			ActivityName:     i.ActivityName,
			LevelAchievement: i.LevelAchievement,
			Levels:           restapi.convertLevelsFromInternal(i.Levels),
			Tags:             restapi.convertTagsFromInternal(i.Tags),
		}
	}

	eItems = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// entriesDay fetches all entries of given date (yyyy-mm-dd)
func (restapi restAPI) entriesDay(w http.ResponseWriter, r *http.Request) {
	year, err := getIntParam(r, `year`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	month, err := getIntParam(r, `month`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	day, err := getIntParam(r, `day`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	fromtime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	precision := time.Hour * 24

	eItems, internalError := restapi.api.EntryGet(context.TODO(), fromtime, precision)

	if internalError != nil {
		panic(internalError)
	}

	if len(eItems) == 0 {
		// No data
		_, _ = io.WriteString(w, `{}`)
		return
	}

	// Convert internal format to JSON API format
	ditems := make(map[string]DTOEntry)

	for _, i := range eItems {
		ditems[i.DateTime.Format(`15:04`)] = DTOEntry{
			Id:               i.Id,
			ActivityName:     i.ActivityName,
			Description:      i.Description,
			LevelAchievement: i.LevelAchievement,
			Levels:           restapi.convertLevelsFromInternal(i.Levels),
			Tags:             restapi.convertTagsFromInternal(i.Tags),
		}
	}

	eItems = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

// entriesLevels gets level names
func (restapi restAPI) entriesLevels(w http.ResponseWriter, r *http.Request) {
	eItems, internalError := restapi.api.EntryLevelsList(context.TODO())

	if internalError != nil {
		panic(internalError)
	}

	if eItems == nil {
		// No data
		_, _ = io.WriteString(w, `[]`)
		return
	}

	// Convert internal format to JSON API format
	var ditems []DTOEntryLevel

	for _, i := range eItems {
		ditems = append(ditems, DTOEntryLevel{
			Id:            i.Id,
			Name:          i.Name,
			ShowByDefault: i.ShowByDefault,
			ShortName:     i.ShortName,
			Worst:         i.Worst,
			AddedAt:       i.AddedAt.Format(`2006-01-02T15:04:05`),
		})
	}

	eItems = nil // Free memory

	b, err := json.Marshal(ditems)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}

func (restapi restAPI) entryLevelUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := getIntParam(r, `id`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	var item DTOEntryLevelUpdate
	if err := readStruct(r.Body, &item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	userErrors, internalError := restapi.api.EntryLevelUpdate(context.TODO(),
		serviceitems.EntryLevelUpdate{
			Id:            id,
			Name:          item.Name,
			ShowByDefault: false,
			ShortName:     item.ShortName,
			Worst:         item.Worst,
			Previous:      false,
		})

	if internalError != nil {
		panic(internalError)
		return
	}

	if userErrors != nil {
		b, err := json.Marshal(userErrors)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(b)
		return
	}

	b, err := json.Marshal(&DTOOK{
		Msg: `ok`,
	})

	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
}
