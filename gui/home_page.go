package gui

import (
	"fmt"
	"log"
	"strings"

	"github.com/Reloyer/rush/dataservice"
	"github.com/Reloyer/rush/gui/widgets"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

var styleCSS string

func ShowHomePage(window *gtk.ApplicationWindow, ds *dataservice.DataService) {

	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(), loadCSS(styleCSS),
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

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
	window.SetTitlebar(gtk.NewBox(gtk.OrientationVertical, 0)) // hide headerbar

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
