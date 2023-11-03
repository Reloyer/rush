package main

import (
	"log"
	"os"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui"
	"github.com/Reloyer/rush/lcu"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {

	lockfile, err := lcu.NewLockfile("/home/Reloyer/Games/wine/prefix/drive_c/Riot Games/League of Legends/lockfile")
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
	userMatches, err := lcu.GetCurrentSummonerMatches(lockfile.Url(), lockfile.TokenEncoded(), 0, 10)

	log.Println(userMatches.Games.Game[4].Participants[0].Stats.Item2)
	ds := dataservice.NewDataService()
	ds.GetHomePageData(userSummonerInfo, userRankedStats)

	app := gtk.NewApplication("com.github.Reloyer.rush", gio.ApplicationFlagsNone)
	app.ConnectActivate((func() { gui.Activite(app, ds) }))

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}

}
