package dataservice

import (
	"fmt"
	"strings"

	"github.com/Reloyer/rush/lcu/datatypes"
	"github.com/Reloyer/rush/utility"
)

type HomeData struct {
	Nickname     string
	ProfileIcon  string
	Level        string
	SoloqIcon    string
	SoloqTier    string
	SoloqDivison string
	SoloqWins    string
	SoloqLoses   string
	SoloqGames   string
	SoloqWR      string
	FlexqIcon    string
	FlexqTier    string
	FlexqDivison string
	FlexqWins    string
	FlexqLoses   string
	FlexqGames   string
	FlexqWR      string
}

type GameData struct {
}
type SettingsData struct {
	Lockfile      string
	WindowsWidth  int
	WindowsHeight int
	language      string
}
type DataService struct {
	Homedata HomeData
	Gamedata GameData
}

func NewDataService() *DataService {
	return &DataService{}
}
func (ds *DataService) GetHomePageData(s datatypes.Summoner, r datatypes.Ranked) {
	assetsDir := "./assets/"
	ds.Homedata.Nickname = s.DisplayName
	ds.Homedata.Level = fmt.Sprintf("%d", s.SummonerLevel)
	ds.Homedata.ProfileIcon = fmt.Sprintf("%sprofileIcons/%d.png", assetsDir, s.ProfileIconId)
	ds.Homedata.SoloqIcon = fmt.Sprintf("%srankedIcons/%s%d.png", assetsDir, strings.ToLower(r.Queues[0].Tier), utility.DivisonToDec(r.Queues[0].Division))
	ds.Homedata.SoloqTier = r.Queues[0].Tier
	ds.Homedata.SoloqDivison = r.Queues[0].Division
	ds.Homedata.SoloqWins = fmt.Sprintf("%d", r.Queues[0].Wins)
	ds.Homedata.SoloqLoses = fmt.Sprintf("%d", r.Queues[0].Losses)
	ds.Homedata.SoloqGames = fmt.Sprintf("%d", (r.Queues[0].Wins + r.Queues[0].Losses))
	ds.Homedata.SoloqWR = fmt.Sprintf("%d", (r.Queues[0].Wins+r.Queues[0].Losses)/r.Queues[0].Wins*10)
	ds.Homedata.FlexqIcon = fmt.Sprintf("%srankedIcons/%s%d.png", assetsDir, strings.ToLower(r.Queues[1].Tier), utility.DivisonToDec(r.Queues[1].Division))
	ds.Homedata.FlexqTier = r.Queues[1].Tier
	ds.Homedata.FlexqDivison = r.Queues[1].Division
	ds.Homedata.FlexqWins = fmt.Sprintf("%d", r.Queues[1].Wins)
	ds.Homedata.FlexqLoses = fmt.Sprintf("%d", r.Queues[1].Losses)
	ds.Homedata.FlexqGames = fmt.Sprintf("%d", (r.Queues[1].Wins + r.Queues[1].Losses))
	ds.Homedata.FlexqWR = fmt.Sprintf("%d", (r.Queues[1].Wins+r.Queues[1].Losses)/r.Queues[1].Wins*10)
}
