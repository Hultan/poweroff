package poweroff

import (
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/poweroff/internal/monitorInfo"
)

type blankWindow struct {
	gtk *gtk.Window
}

// createBlankWindow creates a blank window to block extra screens
// other than the main window.
func createBlankWindow(monitor monitorInfo.MonitorInfo) (*blankWindow, error) {
	aw, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win := &blankWindow{gtk: aw}

	// Set up main window
	win.gtk.SetDecorated(false)
	win.gtk.SetName("window")

	err = makeWindowTransparent(win.gtk)
	if err != nil {
		return nil, err
	}

	// Show the main window
	win.gtk.ShowAll()
	win.gtk.Fullscreen()
	win.gtk.Move(monitor.Left, monitor.Top)
	win.gtk.SetKeepAbove(true)
	win.gtk.SetCanFocus(false)

	return win, nil
}

// closeWindow  destroys the blankWindow
func (w *blankWindow) closeWindow() {
	w.gtk.Close()
	w.gtk.Destroy()
}
