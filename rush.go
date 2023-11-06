package main

import (
	"log"
	"os"
	"sync"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui"
	"github.com/Reloyer/rush/lcu"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {

	lockfile, err := lcu.NewLockfile("/home/Reloyer/Games/league-of-legends/wine/prefix/drive_c/Riot Games/League of Legends/lockfile")
	if err != nil {
		log.Println(err)
	}

	userSummonerInfo, err := lcu.GetCurrentSummoner(lockfile.Url(), lockfile.TokenEncoded())

	if err != nil {
		log.Fatal("Error getting current summoner:", err)
	}

	userRankedStats, err := lcu.GetCurrentRankedStats(lockfile.Url(), lockfile.TokenEncoded())
	if err != nil {
		log.Fatal("Error getting current summoner:", err)
	}
	matchHistory, err := lcu.GetCurrentSummonerMatchHistory(lockfile.Url(), lockfile.TokenEncoded(), 0, 20)

	var wg sync.WaitGroup
	for i := 0; i < len(matchHistory.Games.Game); i++ {
		wg.Add(1)
		i := i
		go func() {
			temp, err := lcu.GetGame(lockfile.Url(), lockfile.TokenEncoded(), matchHistory.Games.Game[i].GameId)
			if err == nil {
				matchHistory.Games.Game[i].ParticipantIdentities = temp.ParticipantIdentities
				matchHistory.Games.Game[i].Participants = temp.Participants
				matchHistory.Games.Game[i].Teams = temp.Teams
			}

			log.Println(i, len(matchHistory.Games.Game[i].Participants))
			defer wg.Done()
		}()
	}
	wg.Wait()
	ds := dataservice.NewDataService()
	ds.GetHomePageData(userSummonerInfo, userRankedStats)
	log.Println("Got Home Page Data")
	ds.GetMatchHistoryData(matchHistory, userSummonerInfo)
	log.Println("Match History Page data")
	app := gtk.NewApplication("com.github.Reloyer.rush", gio.ApplicationFlagsNone)
	app.ConnectActivate((func() { gui.Activite(app, ds) }))

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
