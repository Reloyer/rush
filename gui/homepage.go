package gui

import (
	"fmt"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui/widgets"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func ShowHomePage(w *gtk.ApplicationWindow, ds *dataservice.DataService) {
	ChangeDisplayStyleProvider(gdk.DisplayGetDefault(), "Home", gtk.STYLE_PROVIDER_PRIORITY_APPLICATION, CSSProviders)
	gtk.StyleContextAddProviderForDisplay(gdk.DisplayGetDefault(), CSSProviders["Home"], gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	widgets.CreateNavBar(w, ds, ShowHomePage, ShowMatchHistoryPage)
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

	w.SetChild(grid)
}
