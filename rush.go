package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui"
	"github.com/Reloyer/rush/lcu"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	//configFilePath := "config/config.ini" // specify the path to your config file
	//config, err := config.LoadConfig(configFilePath)
	authToken := "SPDiRasq2QFkmlVg8XPtVg"
	url := "https://127.0.0.1:34037"
	authToken = base64.StdEncoding.EncodeToString([]byte("riot:" + authToken))

	userSummonerInfo, err := lcu.GetCurrentSummoner(url, authToken)

	if err != nil {
		log.Fatal("Error getting current summoner:", err)
	}

	userRankedStats, err := lcu.GetCurrentRankedStats(url, authToken)
	if err != nil {
		log.Fatal("Error getting current summoner:", err)
	}
	//userMatches, err := lcu.GetCurrentSummonerMatches(url, authToken, 0, 10)

	ds := dataservice.NewDataService()
	ds.GetHomePageData(userSummonerInfo, userRankedStats)
	/*userTier := rankstats.Queues[0].Tier
	userDivison := rankstats.Queues[0].Division
	userSoloqWins := rankstats.Queues[0].Wins
	userSoloqLoses := rankstats.Queues[0].Losses

	log.Println(rankstats.Queues[0].Tier, rankstats.Queues[0].Division)
	log.Println(iconURL)
	log.Println(rankURL)
	*/
	app := gtk.NewApplication("com.github.Reloyer.rush", gio.ApplicationFlagsNone)
	app.Connect("activate", func() {
		win := gtk.NewApplicationWindow(app)

		win.SetTitle("Rush")
		win.SetDefaultSize(400, 300)

		// Call the ShowHomePage function here
		gui.ShowHomePage(win, ds)

		win.Show()
	})
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}

}
