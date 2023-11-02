package datatypes

type MatchHistory struct {
	AccountID  int    `json:"accountId"`
	Games      Games  `json:"games"`
	PlatformID string `json:"platformId"`
}

type Games struct {
	GameBeginDate  string       `json:"gameBeginDate"`
	GameCount      int          `json:"gameCount"`
	GameEndDate    string       `json:"gameEndDate"`
	GameIndexBegin int          `json:"gameIndexBegin"`
	GameIndexEnd   int          `json:"gameIndexEnd"`
	Games          []GameDetail `json:"games"`
}

type GameDetail struct {
	GameCreation          int           `json:"gameCreation"`
	GameCreationDate      string        `json:"gameCreationDate"`
	GameDuration          int           `json:"gameDuration"`
	GameID                int           `json:"gameId"`
	GameMode              string        `json:"gameMode"`
	GameType              string        `json:"gameType"`
	GameVersion           string        `json:"gameVersion"`
	MapID                 int           `json:"mapId"`
	ParticipantIdentities []Participant `json:"participantIdentities"`
	Participants          []Participant `json:"participants"`
	PlatformID            string        `json:"platformId"`
	QueueID               int           `json:"queueId"`
	SeasonID              int           `json:"seasonId"`
	Teams                 []TeamDetail  `json:"teams"`
}

type Participant struct {
	ParticipantID int    `json:"participantId"`
	Player        Player `json:"player"`
}

type Player struct {
	AccountID         int    `json:"accountId"`
	CurrentAccountID  int    `json:"currentAccountId"`
	CurrentPlatformID string `json:"currentPlatformId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        int    `json:"summonerId"`
	SummonerName      string `json:"summonerName"`
}

type TeamDetail struct {
	Bans                 []BanDetail `json:"bans"`
	BaronKills           int         `json:"baronKills"`
	DominionVictoryScore int         `json:"dominionVictoryScore"`
	DragonKills          int         `json:"dragonKills"`
	FirstBaron           bool        `json:"firstBaron"`
	FirstBlood           bool        `json:"firstBlood"`
	FirstDragon          bool        `json:"firstDragon"`
	FirstInhibitor       bool        `json:"firstInhibitor"`
	FirstTower           bool        `json:"firstTower"`
	InhibitorKills       int         `json:"inhibitorKills"`
	RiftHeraldKills      int         `json:"riftHeraldKills"`
	TeamID               int         `json:"teamId"`
	TowerKills           int         `json:"towerKills"`
	VilemawKills         int         `json:"vilemawKills"`
	Win                  string      `json:"win"`
}

type BanDetail struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
