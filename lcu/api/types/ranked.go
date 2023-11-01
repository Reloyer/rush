package types

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
	EarnedRegaliaRewardIds                []string `json:"earnedRegaliaRewardIds"`
	HighestPreviousSeasonAchievedDivision string   `json:"highestPreviousSeasonAchievedDivision"`
	HighestPreviousSeasonAchievedTier     string   `json:"highestPreviousSeasonAchievedTier"`
	HighestPreviousSeasonEndDivision      string   `json:"highestPreviousSeasonEndDivision"`
	HighestPreviousSeasonEndTier          string   `json:"highestPreviousSeasonEndTier"`
	HighestRankedEntry                    struct {
		Division                       string `json:"division"`
		IsProvisional                  bool   `json:"isProvisional"`
		LeaguePoints                   int    `json:"leaguePoints"`
		Losses                         int    `json:"losses"`
		MiniSeriesProgress             string `json:"miniSeriesProgress"`
		PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
		PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
		PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
		PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
		ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
		ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
		QueueType                      string `json:"queueType"`
		RatedRating                    int    `json:"ratedRating"`
		RatedTier                      string `json:"ratedTier"`
		Tier                           string `json:"tier"`
		Warnings                       struct {
			DaysUntilDecay                   int  `json:"daysUntilDecay"`
			DemotionWarning                  int  `json:"demotionWarning"`
			DisplayDecayWarning              bool `json:"displayDecayWarning"`
			TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
		} `json:"warnings"`
		Wins int `json:"wins"`
	} `json:"highestRankedEntry"`
	HighestRankedEntrySR struct {
		Division                       string `json:"division"`
		IsProvisional                  bool   `json:"isProvisional"`
		LeaguePoints                   int    `json:"leaguePoints"`
		Losses                         int    `json:"losses"`
		MiniSeriesProgress             string `json:"miniSeriesProgress"`
		PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
		PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
		PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
		PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
		ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
		ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
		QueueType                      string `json:"queueType"`
		RatedRating                    int    `json:"ratedRating"`
		RatedTier                      string `json:"ratedTier"`
		Tier                           string `json:"tier"`
		Warnings                       struct {
			DaysUntilDecay                   int  `json:"daysUntilDecay"`
			DemotionWarning                  int  `json:"demotionWarning"`
			DisplayDecayWarning              bool `json:"displayDecayWarning"`
			TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
		} `json:"warnings"`
		Wins int `json:"wins"`
	} `json:"highestRankedEntrySR"`
	QueueMap struct {
		AdditionalProp1 struct {
			Division                       string `json:"division"`
			IsProvisional                  bool   `json:"isProvisional"`
			LeaguePoints                   int    `json:"leaguePoints"`
			Losses                         int    `json:"losses"`
			MiniSeriesProgress             string `json:"miniSeriesProgress"`
			PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
			PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
			PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
			PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
			ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
			ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
			QueueType                      string `json:"queueType"`
			RatedRating                    int    `json:"ratedRating"`
			RatedTier                      string `json:"ratedTier"`
			Tier                           string `json:"tier"`
			Warnings                       struct {
				DaysUntilDecay                   int  `json:"daysUntilDecay"`
				DemotionWarning                  int  `json:"demotionWarning"`
				DisplayDecayWarning              bool `json:"displayDecayWarning"`
				TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
			} `json:"warnings"`
			Wins int `json:"wins"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Division                       string `json:"division"`
			IsProvisional                  bool   `json:"isProvisional"`
			LeaguePoints                   int    `json:"leaguePoints"`
			Losses                         int    `json:"losses"`
			MiniSeriesProgress             string `json:"miniSeriesProgress"`
			PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
			PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
			PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
			PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
			ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
			ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
			QueueType                      string `json:"queueType"`
			RatedRating                    int    `json:"ratedRating"`
			RatedTier                      string `json:"ratedTier"`
			Tier                           string `json:"tier"`
			Warnings                       struct {
				DaysUntilDecay                   int  `json:"daysUntilDecay"`
				DemotionWarning                  int  `json:"demotionWarning"`
				DisplayDecayWarning              bool `json:"displayDecayWarning"`
				TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
			} `json:"warnings"`
			Wins int `json:"wins"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Division                       string `json:"division"`
			IsProvisional                  bool   `json:"isProvisional"`
			LeaguePoints                   int    `json:"leaguePoints"`
			Losses                         int    `json:"losses"`
			MiniSeriesProgress             string `json:"miniSeriesProgress"`
			PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
			PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
			PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
			PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
			ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
			ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
			QueueType                      string `json:"queueType"`
			RatedRating                    int    `json:"ratedRating"`
			RatedTier                      string `json:"ratedTier"`
			Tier                           string `json:"tier"`
			Warnings                       struct {
				DaysUntilDecay                   int  `json:"daysUntilDecay"`
				DemotionWarning                  int  `json:"demotionWarning"`
				DisplayDecayWarning              bool `json:"displayDecayWarning"`
				TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
			} `json:"warnings"`
			Wins int `json:"wins"`
		} `json:"additionalProp3"`
	} `json:"queueMap"`
	Queues []struct {
		Division                       string `json:"division"`
		IsProvisional                  bool   `json:"isProvisional"`
		LeaguePoints                   int    `json:"leaguePoints"`
		Losses                         int    `json:"losses"`
		MiniSeriesProgress             string `json:"miniSeriesProgress"`
		PreviousSeasonAchievedDivision string `json:"previousSeasonAchievedDivision"`
		PreviousSeasonAchievedTier     string `json:"previousSeasonAchievedTier"`
		PreviousSeasonEndDivision      string `json:"previousSeasonEndDivision"`
		PreviousSeasonEndTier          string `json:"previousSeasonEndTier"`
		ProvisionalGameThreshold       int    `json:"provisionalGameThreshold"`
		ProvisionalGamesRemaining      int    `json:"provisionalGamesRemaining"`
		QueueType                      string `json:"queueType"`
		RatedRating                    int    `json:"ratedRating"`
		RatedTier                      string `json:"ratedTier"`
		Tier                           string `json:"tier"`
		Warnings                       struct {
			DaysUntilDecay                   int  `json:"daysUntilDecay"`
			DemotionWarning                  int  `json:"demotionWarning"`
			DisplayDecayWarning              bool `json:"displayDecayWarning"`
			TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
		} `json:"warnings"`
		Wins int `json:"wins"`
	} `json:"queues"`
	RankedRegaliaLevel int `json:"rankedRegaliaLevel"`
	Seasons            struct {
		AdditionalProp1 struct {
			CurrentSeasonEnd int `json:"currentSeasonEnd"`
			CurrentSeasonId  int `json:"currentSeasonId"`
			NextSeasonStart  int `json:"nextSeasonStart"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			CurrentSeasonEnd int `json:"currentSeasonEnd"`
			CurrentSeasonId  int `json:"currentSeasonId"`
			NextSeasonStart  int `json:"nextSeasonStart"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			CurrentSeasonEnd int `json:"currentSeasonEnd"`
			CurrentSeasonId  int `json:"currentSeasonId"`
			NextSeasonStart  int `json:"nextSeasonStart"`
		} `json:"additionalProp3"`
	} `json:"seasons"`
	SplitsProgress struct {
		AdditionalProp1 int `json:"additionalProp1"`
		AdditionalProp2 int `json:"additionalProp2"`
		AdditionalProp3 int `json:"additionalProp3"`
	} `json:"splitsProgress"`
}
