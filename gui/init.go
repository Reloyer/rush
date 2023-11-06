package gui

import (
	"log"

	"github.com/Reloyer/rush/config"
	"github.com/Reloyer/rush/dataservice"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var CSSProviders = make(map[string]*gtk.CSSProvider)

func Activite(app *gtk.Application, ds *dataservice.DataService) {
	CSSProviders["Home"] = gtk.NewCSSProvider()
	CSSProviders["MatchHistory"] = gtk.NewCSSProvider()

	LoadCSSFromFile(CSSProviders["Home"], "./gui/homepage.css")
	LoadCSSFromFile(CSSProviders["MatchHistory"], "./gui/matchhistorypage.css")

	cfg, err := config.LoadConfig("./config/config.ini")
	if err != nil {
		log.Fatal("fatal error")
	}

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Rush")
	window.SetDefaultSize(cfg.WindowWidth, cfg.WindowHeight)
	window.SetResizable(false)
	window.Show()
	ShowHomePage(window, ds)
}
