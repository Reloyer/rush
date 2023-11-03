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
	GameMode         string
	GameDate         string
	GameResult       string
	GameDuration     string
	UserChampionID   string
	UserChampionIcon string
	ItemIcon0        string
	ItemIcon1        string
	ItemIcon2        string
	ItemIcon3        string
	ItemIcon4        string
	ItemIcon5        string
	ItemIcon6        string
}
type SettingsData struct {
	Lockfile      string
	WindowsWidth  int
	WindowsHeight int
	Language      string
}
type DataService struct {
	Homedata HomeData
	Gamedata GameData
}

func NewDataService() *DataService {
	return &DataService{}
}
func (ds *DataService) GetHomePageData(s datatypes.Summoner, r datatypes.Ranked) {
	assetsDir := "./assets"
	profileIconsDir := assetsDir + "/profile-icons"
	rankIconsDir := assetsDir + "/rank-icons"

	hd := &ds.Homedata
	hd.Nickname = s.DisplayName
	hd.Level = fmt.Sprintf("%d", s.SummonerLevel)
	hd.ProfileIcon = fmt.Sprintf("%s/%d.png", profileIconsDir, s.ProfileIconId)

	hd.SoloqIcon = fmt.Sprintf("%s%s/%d.png", rankIconsDir, strings.ToLower(r.Queues[0].Tier), utility.DivisonToDec(r.Queues[0].Division))
	hd.SoloqTier = r.Queues[0].Tier
	hd.SoloqDivison = r.Queues[0].Division
	hd.SoloqWins = fmt.Sprintf("%d", r.Queues[0].Wins)
	hd.SoloqLoses = fmt.Sprintf("%d", r.Queues[0].Losses)
	hd.SoloqGames = fmt.Sprintf("%d", (r.Queues[0].Wins + r.Queues[0].Losses))
	hd.SoloqWR = fmt.Sprintf("%d", (r.Queues[0].Wins+r.Queues[0].Losses)/r.Queues[0].Wins*10)

	hd.FlexqIcon = fmt.Sprintf("%s%s/%d.png", rankIconsDir, strings.ToLower(r.Queues[1].Tier), utility.DivisonToDec(r.Queues[1].Division))
	hd.FlexqTier = r.Queues[1].Tier
	hd.FlexqDivison = r.Queues[1].Division
	hd.FlexqWins = fmt.Sprintf("%d", r.Queues[1].Wins)
	hd.FlexqLoses = fmt.Sprintf("%d", r.Queues[1].Losses)
	hd.FlexqGames = fmt.Sprintf("%d", (r.Queues[1].Wins + r.Queues[1].Losses))
	hd.FlexqWR = fmt.Sprintf("%d", (r.Queues[1].Wins+r.Queues[1].Losses)/r.Queues[1].Wins*10)
}
func (ds *DataService) GetGameData(g datatypes.Game, s datatypes.Summoner) {
	assetsDir := "./assets"
	championIcons := assetsDir + "/champion-icons"
	itemsIcons := assetsDir + "/items-icons"

	i := 0
	userIndex := 0
	for i < len(g.Participants) {
		if g.Participants[i].ParticipantId == s.SummonerId {
			userIndex = i
		}
		i++
	}
	gd := &ds.Gamedata
	gd.GameMode = g.GameMode
	gd.GameDate = g.GameCreationDate
	gd.GameDuration = fmt.Sprintf("%.2f", float64(g.GameDuration)/60)
	if g.Participants[userIndex].Stats.Win == false {
		gd.GameResult = "Lose"
	} else {
		gd.GameResult = "Win"
	}
	gd.UserChampionID = fmt.Sprintf("%d", g.Participants[userIndex].ChampionId)
	gd.UserChampionIcon = fmt.Sprintf("%s/%d.png", championIcons, g.Participants[userIndex].ChampionId)

	gd.ItemIcon0 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item0)
	gd.ItemIcon1 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item1)
	gd.ItemIcon2 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item2)
	gd.ItemIcon3 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item3)
	gd.ItemIcon4 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item4)
	gd.ItemIcon5 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item5)
	gd.ItemIcon6 = fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userIndex].Stats.Item6)
}
