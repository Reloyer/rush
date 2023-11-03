package gui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Reloyer/rush/config"
	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui/widgets"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type HomePage struct {
}

var styleCSS string

func Activite(app *gtk.Application, ds *dataservice.DataService) {
	b, err := os.ReadFile("./gui/homepage.css") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	styleCSS = string(b)
	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(), loadCSS(styleCSS),
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	grid := gtk.NewGrid()
	grid.SetColumnSpacing(6)
	grid.SetRowSpacing(6)
	img_usrIcon := gtk.NewImageFromFile(ds.Homedata.ProfileIcon)
	img_usrIcon.SetSizeRequest(200, 200)

	l_usrNick := gtk.NewLabel(ds.Homedata.Nickname)
	l_usrNick.AddCSSClass("title")
	l_usrLevel := gtk.NewLabel(ds.Homedata.Level)
	l_usrLevel.AddCSSClass("subtitle")

	grid.Attach(img_usrIcon, 0, 0, 2, 2)
	grid.Attach(l_usrNick, 3, 0, 1, 1)
	grid.Attach(l_usrLevel, 0, 1, 2, 1)

	img_soloqRankImage := gtk.NewImageFromFile(ds.Homedata.SoloqIcon)
	img_soloqRankImage.SetSizeRequest(200, 200)
	l_soloqstats := gtk.NewLabel(fmt.Sprintf("%s %s Games:%s \nWinRate:%s Wins:%s Loses:%s ", ds.Homedata.SoloqTier, ds.Homedata.SoloqDivison, ds.Homedata.SoloqGames,
		ds.Homedata.SoloqWR, ds.Homedata.SoloqWins, ds.Homedata.SoloqLoses))
	l_soloqstats.AddCSSClass("soloq-stats")

	grid.Attach(img_soloqRankImage, 0, 4, 2, 2)
	grid.Attach(l_soloqstats, 2, 4, 2, 2)

	flexqRankImage := gtk.NewImageFromFile(ds.Homedata.FlexqIcon)
	flexqRankImage.SetSizeRequest(200, 200)
	l_flexqstats := gtk.NewLabel(fmt.Sprintf("%s %s Games:%s \nWinRate:%s Wins:%s Loses:%s ", ds.Homedata.FlexqTier, ds.Homedata.FlexqDivison, ds.Homedata.FlexqGames,
		ds.Homedata.FlexqWR, ds.Homedata.FlexqWins, ds.Homedata.FlexqLoses))
	l_flexqstats.AddCSSClass("soloq-stats")

	grid.Attach(flexqRankImage, 0, 6, 2, 2)
	grid.Attach(l_flexqstats, 2, 6, 2, 2)

	match := widgets.NewGame()

	cfg, err := config.LoadConfig("./config/config.ini")
	if err != nil {
		log.Fatal("fatal error")
	}

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Rush")
	window.SetChild(match)
	window.SetDefaultSize(cfg.WindowWidth, cfg.WindowHeight)
	window.Show()
}

func loadCSS(content string) *gtk.CSSProvider {
	prov := gtk.NewCSSProvider()
	prov.ConnectParsingError(func(sec *gtk.CSSSection, err error) {
		// Optional line parsing routine.
		loc := sec.StartLocation()
		lines := strings.Split(content, "\n")
		log.Printf("CSS error (%v) at line: %q", err, lines[loc.Lines()])
	})
	prov.LoadFromData(content)
	return prov
}
