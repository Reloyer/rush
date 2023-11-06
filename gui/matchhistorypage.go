package gui

import (
	"log"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui/widgets"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func ShowMatchHistoryPage(w *gtk.ApplicationWindow, ds *dataservice.DataService) {
	//display := gdk.DisplayGetDefault()
	ChangeDisplayStyleProvider(gdk.DisplayGetDefault(), "MatchHistory", gtk.STYLE_PROVIDER_PRIORITY_APPLICATION, CSSProviders)
	widgets.CreateNavBar(w, ds, ShowHomePage, ShowMatchHistoryPage)
	main_box := gtk.NewBox(gtk.OrientationVertical, 8)
	scrolledwindow := gtk.NewScrolledWindow()
	scrolledwindow.SetVExpand(true)
	box := gtk.NewBox(gtk.OrientationVertical, 8)
	log.Println("Match History box added")
	mhd := ds.MatchHistorydata
	for i := 0; i < mhd.MatchHistory.Games.GameCount; i++ {
		b_game := widgets.NewGame(mhd.Gamedata[i])
		if mhd.Gamedata[i].GameResult == "VICTORY" {
			b_game.AddCSSClass("gamebox-won")
			log.Println(b_game.StyleContext())
		} else if mhd.Gamedata[i].GameResult == "DEFEAT" {
			b_game.AddCSSClass("gamebox-lost")
		} else {
			b_game.AddCSSClass("gamebox-unkown")
		}
		box.Append(b_game)
	}
	scrolledwindow.SetChild(box)
	main_box.Append(scrolledwindow)
	// Set the main_box as the child of scrolledWindow
	w.SetChild(main_box)
}
