package restapi

// JSON formats for REST API

type DTOEntryAdd struct {
	Activity         string         `json:"activity"`
	Description      string         `json:"description"`
	LevelAchievement int            `json:"achievement"`
	Levels           map[string]int `json:"levels"`
	Tags             []string       `json:"tags"`
}

type DTOEntry struct {
	Id               int            `json:"id"`
	ActivityName     string         `json:"activity"`
	Description      string         `json:"description"`
	LevelAchievement int            `json:"achievement"`
	Levels           map[string]int `json:"levels"`
	Tags             []string       `json:"tags"`
}

type DTOOK struct {
	Msg string `json:"msg"`
}

type DTONewId struct {
	Id int `json:"id"`
}

type DTOTask struct {
	Id                         int     `json:"id"`
	AddedAt                    string  `json:"added_at"`
	CompletedAt                *string `json:"completed_at"`
	Title                      string  `json:"title"`
	Description                string  `json:"description"`
	ReoccurringTaskReferenceId *int    `json:"refid,omitempty"`
}

type DTOTimerAdd struct {
	Title   string `json:"title"`
	Seconds uint64 `json:"seconds"`
}

type DTOTimer struct {
	Id        uint64  `json:"id"`
	Seconds   float64 `json:"s"`
	Formatted string  `json:"f"`
}

type DTOEntryLevelUpdate struct {
	Name          string `json:"name"`
	ShortName     string `json:"key"`
	ShowByDefault bool   `json:"show"`
	Worst         string `json:"worst"`
}

type DTOEntryLevel struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ShortName     string `json:"key"`
	ShowByDefault bool   `json:"show"`
	Worst         string `json:"worst"`
	AddedAt       string `json:"added_at"`
}

type DTOReoccurringTaskAdd struct {
	Title   string `json:"title"`
	Seconds int    `json:"s"`
}

type DTOReoccurringTask struct {
	Id      int    `json:"id"`
	AddedAt string `json:"added_at"`
	Title   string `json:"title"`
}

type DTOWeightAdd struct {
	Weight float32 `json:"weight"`
}

type DTOWeight struct {
	Weight float32 `json:"weight"`
	Added  string  `json:"added"`
	Id     int     `json:"id"`
}

type DTOHeightAdd struct {
	Height float32 `json:"height"`
}

type DTOHeight struct {
	Height float32 `json:"height"`
	Added  string  `json:"added"`
	Id     int     `json:"id"`
}

type DTOTestMADRSAnswers struct {
	Answer1  int `json:"a1"`
	Answer2  int `json:"a2"`
	Answer3  int `json:"a3"`
	Answer4  int `json:"a4"`
	Answer5  int `json:"a5"`
	Answer6  int `json:"a6"`
	Answer7  int `json:"a7"`
	Answer8  int `json:"a8"`
	Answer9  int `json:"a9"`
	Answer10 int `json:"a10"`
}

type DTOTag struct {
	Id        int    `json:"id"`
	AddedAt   string `json:"added_at"`
	Name      string `json:"name"`
	ShortName string `json:"shortname"`
}

type DTOResolutionsUpdate struct {
	EntityId     int    `json:"entityid"`
	Name         string `json:"name"`
	DecisionDate string `json:"decisiondate"`
	SentDate     string `json:"sentdate"`
	StartDate    string `json:"startdate"`
	EndDate      string `json:"enddate"`
}

type DTOResolutions struct {
	Id           int      `json:"id"`
	EntityId     int      `json:"entityid"`
	AddedAt      string   `json:"added_at"`
	Name         string   `json:"name"`
	DecisionDate string   `json:"decisiondate"`
	SentDate     string   `json:"sentdate"`
	StartDate    string   `json:"startdate"`
	EndDate      string   `json:"enddate"`
	Files        []string `json:"files"`
}

type DTOResolutionEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
