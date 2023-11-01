package main

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"

	"github.com/Reloyer/rush/gui"
	"github.com/Reloyer/rush/lcu"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	authToken := "UI9cT03SJ-BLs2Sokk-hzg"
	url := "https://127.0.0.1:37849"
	authToken = base64.StdEncoding.EncodeToString([]byte("riot:" + authToken))

	uSumInfo, err := lcu.GetCurrentSummoner(url, authToken)

	if err != nil {
		log.Fatal("Error getting current summoner:", err)
	}

	uRankStats, err := lcu.GetCurrentRankedStats(url, authToken)

	/*userTier := rankstats.Queues[0].Tier
	userDivison := rankstats.Queues[0].Division
	userSoloqWins := rankstats.Queues[0].Wins
	userSoloqLoses := rankstats.Queues[0].Losses

	log.Println(rankstats.Queues[0].Tier, rankstats.Queues[0].Division)
	log.Println(iconURL)
	log.Println(rankURL)
	*/
	app := gtk.NewApplication("com.github.Reloyer.rush", gio.ApplicationFlagsNone)
	app.ConnectActivate((func() { gui.Activite(app, uSumInfo, uRankStats) }))

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}

}

// waitForExitSignal blocks the main goroutine until an exit signal is received
func waitForExitSignal() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Press Enter to exit...")
	scanner.Scan()
}
