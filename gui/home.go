package gui

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Reloyer/rush/lcu/api/types"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

var styleCSS string

func Activite(app *gtk.Application, sumsinfo types.Summoner, rankedstats types.Ranked) {

	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(), loadCSS(styleCSS),
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	title := gtk.NewLabel(fmt.Sprintf("%s", sumsinfo.DisplayName))
	title.AddCSSClass("title")
	title.SetWrap(true)
	title.SetWrapMode(pango.WrapWordChar)
	title.SetXAlign(0)
	title.SetYAlign(0)

	log.Println(fmt.Sprintf("assets/profileIcons/%i.jpg", sumsinfo.ProfileIconId))
	image := gtk.NewImageFromFile(("assets/profileIcons/" + strconv.Itoa(sumsinfo.ProfileIconId) + ".jpg")) // Replace with your image path

	soloqRankImage := gtk.NewImageFromFile(fmt.Sprintf("assets/rankedIcons/%s%d.png", strings.ToLower(rankedstats.Queues[0].Tier), rankedstats.Queues[0].Division))
	flexRankImage := gtk.NewImageFromFile(fmt.Sprintf("assets/rankedIcons/%s%d.png", strings.ToLower(rankedstats.Queues[1].Tier), rankedstats.Queues[1].Division))

	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.Append(title)
	box.Append(image) // Add the image to the box
	box.Append(soloqRankImage)
	box.Append(flexRankImage)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Anti-browser")
	window.SetChild(box)
	window.SetTitlebar(gtk.NewBox(gtk.OrientationVertical, 0)) // hide headerbar
	window.SetDefaultSize(400, 300)
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
