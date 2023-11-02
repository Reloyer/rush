package datatypes

var DivToDec = map[string]int{
	"I":   1,
	"II":  2,
	"III": 3,
	"IV":  4,
}

const (
	SOLOQ = 0
	FlEXQ = 1
)

type Ranked struct {
	EarnedRegaliaRewardIds                []string                 `json:"earnedRegaliaRewardIds"`
	HighestPreviousSeasonAchievedDivision string                   `json:"highestPreviousSeasonAchievedDivision"`
	HighestPreviousSeasonAchievedTier     string                   `json:"highestPreviousSeasonAchievedTier"`
	HighestPreviousSeasonEndDivision      string                   `json:"highestPreviousSeasonEndDivision"`
	HighestPreviousSeasonEndTier          string                   `json:"highestPreviousSeasonEndTier"`
	HighestRankedEntry                    RankedEntry              `json:"highestRankedEntry"`
	HighestRankedEntrySR                  RankedEntrySR            `json:"highestRankedEntrySR"`
	QueueMap                              map[string]QueueEntry    `json:"queueMap"`
	Queues                                []QueueEntry             `json:"queues"`
	RankedRegaliaLevel                    int                      `json:"rankedRegaliaLevel"`
	Seasons                               map[string]SeasonDetails `json:"seasons"`
	SplitsProgress                        map[string]int           `json:"splitsProgress"`
}

type RankedEntry struct {
	Division                       string   `json:"division"`
	IsProvisional                  bool     `json:"isProvisional"`
	LeaguePoints                   int      `json:"leaguePoints"`
	Losses                         int      `json:"losses"`
	MiniSeriesProgress             string   `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string   `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string   `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string   `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string   `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int      `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int      `json:"provisionalGamesRemaining"`
	QueueType                      string   `json:"queueType"`
	RatedRating                    int      `json:"ratedRating"`
	RatedTier                      string   `json:"ratedTier"`
	Tier                           string   `json:"tier"`
	Warnings                       Warnings `json:"warnings"`
	Wins                           int      `json:"wins"`
}
type RankedEntrySR struct {
	Division                       string   `json:"division"`
	IsProvisional                  bool     `json:"isProvisional"`
	LeaguePoints                   int      `json:"leaguePoints"`
	Losses                         int      `json:"losses"`
	MiniSeriesProgress             string   `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string   `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string   `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string   `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string   `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int      `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int      `json:"provisionalGamesRemaining"`
	QueueType                      string   `json:"queueType"`
	RatedRating                    int      `json:"ratedRating"`
	RatedTier                      string   `json:"ratedTier"`
	Tier                           string   `json:"tier"`
	Warnings                       Warnings `json:"warnings"`
	Wins                           int      `json:"wins"`
}
type HighestRankedEntrySR struct {
	Division                       string   `json:"division"`
	IsProvisional                  bool     `json:"isProvisional"`
	LeaguePoints                   int      `json:"leaguePoints"`
	Losses                         int      `json:"losses"`
	MiniSeriesProgress             string   `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string   `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string   `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string   `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string   `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int      `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int      `json:"provisionalGamesRemaining"`
	QueueType                      string   `json:"queueType"`
	RatedRating                    int      `json:"ratedRating"`
	RatedTier                      string   `json:"ratedTier"`
	Tier                           string   `json:"tier"`
	Warnings                       Warnings `json:"warnings"`
	Wins                           int      `json:"wins"`
}

// Warnings represents the warnings data within HighestRankedEntrySR.
type Warnings struct {
	DaysUntilDecay                   int  `json:"daysUntilDecay"`
	DemotionWarning                  int  `json:"demotionWarning"`
	DisplayDecayWarning              bool `json:"displayDecayWarning"`
	TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
}

type QueueEntry struct {
	Division                       string   `json:"division"`
	IsProvisional                  bool     `json:"isProvisional"`
	LeaguePoints                   int      `json:"leaguePoints"`
	Losses                         int      `json:"losses"`
	MiniSeriesProgress             string   `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string   `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string   `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string   `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string   `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int      `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int      `json:"provisionalGamesRemaining"`
	QueueType                      string   `json:"queueType"`
	RatedRating                    int      `json:"ratedRating"`
	RatedTier                      string   `json:"ratedTier"`
	Tier                           string   `json:"tier"`
	Warnings                       Warnings `json:"warnings"`
	Wins                           int      `json:"wins"`
}

type SeasonDetails struct {
	CurrentSeasonEnd int `json:"currentSeasonEnd"`
	CurrentSeasonId  int `json:"currentSeasonId"`
	NextSeasonStart  int `json:"nextSeasonStart"`
}
