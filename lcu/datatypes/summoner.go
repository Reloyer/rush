package datatypes

type Summoner struct {
	AccountId                   int          `json:"accountId"`
	DisplayName                 string       `json:"displayName"`
	InternalName                string       `json:"internalName"`
	NameChangeFlag              bool         `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int          `json:"percentCompleteForNextLevel"`
	Privacy                     string       `json:"privacy"`
	ProfileIconId               int          `json:"profileIconId"`
	Puuid                       string       `json:"puuid"`
	RerollPoints                RerollPoints `json:"rerollPoints"`
	SummonerId                  int          `json:"summonerId"`
	SummonerLevel               int          `json:"summonerLevel"`
	Unnamed                     bool         `json:"unnamed"`
	XpSinceLastLevel            int          `json:"xpSinceLastLevel"`
	XpUntilNextLevel            int          `json:"xpUntilNextLevel"`
}

// RerollPoints represents the structure of reroll points for a summoner.
type RerollPoints struct {
	CurrentPoints    int `json:"currentPoints"`
	MaxRolls         int `json:"maxRolls"`
	NumberOfRolls    int `json:"numberOfRolls"`
	PointsCostToRoll int `json:"pointsCostToRoll"`
	PointsToReroll   int `json:"pointsToReroll"`
}
