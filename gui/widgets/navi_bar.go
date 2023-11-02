package widgets

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// CreateNavBar creates a navigation bar.
func CreateNavBar(window *gtk.ApplicationWindow, showHomePage func(), showSettingsPage func()) {
	headerBar := gtk.NewHeaderBar()

	homePageButton := gtk.NewButton()
	homePageButton.SetLabel("Home Page")
	homePageButton.Connect("clicked", func() {
		showHomePage() // Call function to switch to page 1
	})

	settingsPageButton := gtk.NewButton()
	settingsPageButton.SetLabel("Page 2")
	settingsPageButton.Connect("clicked", func() {
		showSettingsPage() // Call function to switch to page 2
	})

	headerBar.PackStart(homePageButton)
	headerBar.PackStart(settingsPageButton)

	window.SetTitlebar(headerBar)
}
