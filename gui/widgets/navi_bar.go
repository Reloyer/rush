package widgets

import (
	"github.com/Reloyer/rush/dataservice"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// CreateNavBar creates a navigation bar.
func CreateNavBar(w *gtk.ApplicationWindow, ds *dataservice.DataService, showHomePage func(*gtk.ApplicationWindow, *dataservice.DataService), showMatchHistoryPage func(*gtk.ApplicationWindow, *dataservice.DataService)) {
	headerBar := gtk.NewHeaderBar()
	homePageButton := gtk.NewButton()
	homePageButton.SetLabel("Home Page")
	homePageButton.Connect("clicked", func() {
		showHomePage(w, ds) // Call function to switch to page HomePage
	})

	matchHistoryPageButton := gtk.NewButton()
	matchHistoryPageButton.SetLabel("MatchHistory")
	matchHistoryPageButton.Connect("clicked", func() {
		showMatchHistoryPage(w, ds) // Call function to switch to MatchHistoryPage
	})

	headerBar.PackStart(homePageButton)
	headerBar.PackStart(matchHistoryPageButton)

	w.SetTitlebar(headerBar)
}
