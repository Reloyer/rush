package gui

import (
	"log"
	"os"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func LoadCSSFromFile(prov *gtk.CSSProvider, path string) (err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error while Opening CSS file: ", err)
		return err
	}
	LoadCSSFromString(prov, string(content))
	return nil
}

func LoadCSSFromString(prov *gtk.CSSProvider, content string) {
	prov.ConnectParsingError(func(sec *gtk.CSSSection, err error) {
		loc := sec.StartLocation()
		lines := strings.Split(content, "\n")
		log.Printf("CSS error (%v) at line: %q", err, lines[loc.Lines()])
	})
	prov.LoadFromData(content)
}

func WidgetApplyStyleProvider(widget gtk.Widget, p *gtk.CSSProvider) {
	c := widget.StyleContext()
	c.AddProvider(p, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func WidgetRemoveStyleProviderFrom(w *gtk.Widget, p *gtk.CSSProvider) {
	context := w.StyleContext()
	context.RemoveProvider(p)
}

func RemoveAllStlyeProvidersFromDisplay(display *gdk.Display, providers map[string]*gtk.CSSProvider) {
	for _, provider := range providers {
		gtk.StyleContextRemoveProviderForDisplay(display, provider)
	}
}
func ChangeDisplayStyleProvider(display *gdk.Display, name string, priority uint, providers map[string]*gtk.CSSProvider) {
	provider, ok := providers[name]
	if ok {
		RemoveAllStlyeProvidersFromDisplay(display, providers)
		gtk.StyleContextAddProviderForDisplay(display, provider, priority)
		log.Println("Success: Changed StyleProvider to ", name)
	} else {
		log.Println("ERROR: ChagneDisplayStyleProvider() error: CSS Provider is not it CSSProviders map")
	}

}
