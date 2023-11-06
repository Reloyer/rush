package dataservice

import (
	"fmt"
	"log"
	"strings"
	"time"

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
type Player struct {
	Nickname     string
	ChampionIcon string
}
type GameData struct {
	UserInedx          int
	GameType           string
	TimeSince          string
	GameResult         string
	GameDuration       string
	UserNickName       string
	UserChampionID     string
	UserChampionIcon   string
	UserKDA            string
	UserKDAPerfect     bool
	UserKDASuperior    bool
	UserKDAGood        bool
	UserKDANormal      bool
	UserKDABad         bool
	UserRatio          string
	ItemsIcons         []string
	Perk0Icon          string
	PerkSubStyleIcon   string
	SummonerSpell1Icon string
	SummonerSpell2Icon string
	NumberOfPlayer     int
	Team1              []Player
	Team2              []Player
}
type MatchHistoryData struct {
	MatchHistory datatypes.MatchHistory
	Gamedata     []GameData
	UserNickName string
}

var gameType = map[int]string{
	0:   "Custom Game",
	400: "Normal (Draft Pick)",
	420: "Ranked Solo",
	430: "Normal (Blind Pick)",
	440: "Ranked Flex",
	450: "Aram",
	700: "Clash",
	720: "Clash Aram",
	830: "Co-op vs AI (Intro)",
	840: "Co-op vs AI (Beginner)",
	850: "Co-op vs AI (Advanced)",
}

type SettingsData struct {
	Lockfile      string
	WindowsWidth  int
	WindowsHeight int
	Language      string
}
type DataService struct {
	Homedata         HomeData
	MatchHistorydata MatchHistoryData
}

func NewDataService() *DataService {
	return &DataService{}
}
func (ds *DataService) GetHomePageData(s datatypes.Summoner, r datatypes.Ranked) {
	assetsDir := "./assets"
	profileIconsDir := assetsDir + "/profile-icons"
	rankIconsDir := assetsDir + "/ranked-icons"

	hd := &ds.Homedata
	hd.Nickname = s.DisplayName
	hd.Level = fmt.Sprintf("%d", s.SummonerLevel)
	hd.ProfileIcon = fmt.Sprintf("%s/%d.jpg", profileIconsDir, s.ProfileIconId)

	hd.SoloqIcon = fmt.Sprintf("%s/%s%d.png", rankIconsDir, strings.ToLower(r.Queues[0].Tier), utility.DivisonToDec(r.Queues[0].Division))
	hd.SoloqTier = r.Queues[0].Tier
	hd.SoloqDivison = r.Queues[0].Division
	hd.SoloqWins = fmt.Sprintf("%d", r.Queues[0].Wins)
	hd.SoloqLoses = fmt.Sprintf("%d", r.Queues[0].Losses)
	hd.SoloqGames = fmt.Sprintf("%d", (r.Queues[0].Wins + r.Queues[0].Losses))
	hd.SoloqWR = fmt.Sprintf("%d", (r.Queues[0].Wins+r.Queues[0].Losses)/r.Queues[0].Wins*10)

	hd.FlexqIcon = fmt.Sprintf("%s/%s%d.png", rankIconsDir, strings.ToLower(r.Queues[1].Tier), utility.DivisonToDec(r.Queues[1].Division))
	hd.FlexqTier = r.Queues[1].Tier
	hd.FlexqDivison = r.Queues[1].Division
	hd.FlexqWins = fmt.Sprintf("%d", r.Queues[1].Wins)
	hd.FlexqLoses = fmt.Sprintf("%d", r.Queues[1].Losses)
	hd.FlexqGames = fmt.Sprintf("%d", (r.Queues[1].Wins + r.Queues[1].Losses))
	hd.FlexqWR = fmt.Sprintf("%d", (r.Queues[1].Wins+r.Queues[1].Losses)/r.Queues[1].Wins*10)
}
func (ds *DataService) GetGameData(g datatypes.Game, s datatypes.Summoner) GameData {
	assetsDir := "./assets"
	championIcons := assetsDir + "/champion-icons"
	itemsIcons := assetsDir + "/items-icons"
	perkIcons := assetsDir + "/perk-icons"
	summonerIcons := assetsDir + "/summoners-icons"

	userId := -1

	for i, participant := range g.ParticipantIdentities {
		if participant.Player.SummonerName == s.DisplayName {
			userId = i
		}
	}
	if userId == -1 {
		log.Println("ERROR GetGameData error: didnt found User dispaly name in game data")
		return GameData{}
	}

	gd := GameData{}
	gd.UserInedx = userId
	gd.UserNickName = s.DisplayName
	gd.UserChampionID = fmt.Sprintf("%d", g.Participants[userId].ChampionId)
	gd.UserChampionIcon = fmt.Sprintf("%s/%d.png", championIcons, g.Participants[userId].ChampionId)
	gd.UserKDA = fmt.Sprintf("%d / %d / %d", g.Participants[userId].Stats.Kills, g.Participants[userId].Stats.Deaths, g.Participants[userId].Stats.Assists)
	gd.UserRatio = fmt.Sprintf("%.2f KDA", (float64(g.Participants[userId].Stats.Kills)+float64(g.Participants[userId].Stats.Assists))/float64(g.Participants[userId].Stats.Deaths))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item0))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item1))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item2))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item3))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item4))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item5))
	gd.ItemsIcons = append(gd.ItemsIcons, fmt.Sprintf("%s/%d.png", itemsIcons, g.Participants[userId].Stats.Item6))
	gd.SummonerSpell1Icon = fmt.Sprintf("%s/%d.png", summonerIcons, g.Participants[userId].Spell1Id)
	gd.SummonerSpell2Icon = fmt.Sprintf("%s/%d.png", summonerIcons, g.Participants[userId].Spell2Id)
	gd.Perk0Icon = fmt.Sprintf("%s/%d.png", perkIcons, g.Participants[userId].Stats.Perk0)
	gd.PerkSubStyleIcon = fmt.Sprintf("%s/%d.png", perkIcons, g.Participants[userId].Stats.PerkSubStyle)
	gd.NumberOfPlayer = len(g.Participants)
	gd.GameDuration = fmt.Sprintf("%dm %ds", int(g.GameDuration)/60, int(g.GameDuration)%60)

	gd.GameType = gameType[g.QueueId]
	if gd.GameType == "" {
		gd.GameType = "Unknown"
	}

	if g.Participants[userId].Stats.Win == false {
		gd.GameResult = "VICTORY"
	} else {
		gd.GameResult = "DEFEAT"
	}

	userratio := (float64(g.Participants[userId].Stats.Kills) + float64(g.Participants[userId].Stats.Assists)) / float64(g.Participants[userId].Stats.Deaths)
	if userratio >= 4.5 {
		gd.UserKDASuperior = true
	} else if userratio >= 2.5 {
		gd.UserKDAGood = true
	} else if userratio >= 1.5 {
		gd.UserKDANormal = true
	} else if g.Participants[userId].Stats.Deaths == 0 {
		gd.UserKDANormal = true
	} else {
		gd.UserKDABad = true
	}

	currentTime := time.Now()
	timeElapsed := currentTime.Sub(g.GameCreationDate)
	if timeElapsed.Hours() > 24 {
		gd.TimeSince = fmt.Sprintf("%d days ago", int(timeElapsed.Hours()/24))
	} else if timeElapsed.Hours() >= 1 {
		gd.TimeSince = fmt.Sprintf("%d hours ago", int(timeElapsed.Hours()))
	} else {
		gd.TimeSince = fmt.Sprintf("%d minutes ago", int(timeElapsed.Minutes()))
	}

	for i, participant := range g.Participants {
		p := Player{
			Nickname:     g.ParticipantIdentities[i].Player.SummonerName,                  // Replace with the appropriate field from your data
			ChampionIcon: fmt.Sprintf("%s/%d.png", championIcons, participant.ChampionId), // Replace with the appropriate field from your data
		}
		if g.Teams[0].TeamId == participant.TeamId {
			gd.Team1 = append(gd.Team1, p)
		} else {
			gd.Team2 = append(gd.Team2, p)
		}
	}

	return gd
}
func (ds *DataService) GetMatchHistoryData(mh datatypes.MatchHistory, s datatypes.Summoner) {
	mhd := &ds.MatchHistorydata
	mhd.UserNickName = s.DisplayName
	mhd.MatchHistory = mh
	for i := 0; i < mhd.MatchHistory.Games.GameCount; i++ {
		mhd.Gamedata = append(mhd.Gamedata, ds.GetGameData(mh.Games.Game[i], s))
	}
}
