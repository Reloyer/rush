package gui

import (
	"fmt"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui/widgets"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// showSettingsPage displays the settings page.
func ShowSettingsPage(window *gtk.ApplicationWindow, ds *dataservice.DataService) {
	showHomePage := func() {
		ShowHomePage(window, ds)
	}
	showSettingsPage := func() {
		ShowSettingsPage(window, ds)
	}
	widgets.CreateNavBar(window, showHomePage, showSettingsPage)

	title := gtk.NewLabel(fmt.Sprintf("%s", ds.Homedata.Nickname))
	title.AddCSSClass("title")
	title.SetWrap(true)
	title.SetWrapMode(pango.WrapWordChar)
	title.SetXAlign(0)
	title.SetYAlign(0)

	profileIconImage := gtk.NewImageFromFile(ds.Homedata.ProfileIcon)

	soloqRankImage := gtk.NewImageFromFile(ds.Homedata.SoloqIcon)
	flexRankImage := gtk.NewImageFromFile(ds.Homedata.FlexqIcon)

	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.Append(title)
	box.Append(profileIconImage) // Add the image to the box
	box.Append(soloqRankImage)
	box.Append(flexRankImage)

	window.SetTitle("Rush")
	window.SetChild(box)
	window.SetTitlebar(gtk.NewBox(gtk.OrientationVertical, 0))
}
