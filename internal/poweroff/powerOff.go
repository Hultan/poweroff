package poweroff

import (
	"errors"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/poweroff/internal/monitorInfo"
)

const applicationTitle = "poweroff"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type PowerOff struct {
}

// StartApplication : Starts the PowerOff application
func (p *PowerOff) StartApplication(app *gtk.Application) (err error) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Get information about the computers monitors
	info := monitorInfo.NewSoftMonitorInfo()
	monitors, err := info.GetMonitorInfo()
	if err != nil {
		return err
	}

	_, err = createMainWindow(app, monitors)
	if err != nil {
		return err
	}

	gtk.Main()

	return nil
}

func makeWindowTransparent(win interface{}) (err error) {
	var (
		cssProv *gtk.CssProvider
		screen  *gdk.Screen
		visual  *gdk.Visual
	)

	if cssProv, err = gtk.CssProviderNew(); err != nil {
		return
	}
	if err = cssProv.LoadFromData(css); err != nil {
		return
	}
	if screen, err = gdk.ScreenGetDefault(); err != nil {
		return
	}

	gtk.AddProviderForScreen(screen, cssProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	switch w := win.(type) {
	case *gtk.ApplicationWindow:
		if visual, err = w.GetScreen().GetRGBAVisual(); err == nil {
			w.SetVisual(visual)
		}
	case *gtk.Window:
		if visual, err = w.GetScreen().GetRGBAVisual(); err == nil {
			w.SetVisual(visual)
		}
	default:
		return errors.New("invalid window type in makeWindowTransparent()")
	}
	return
}
