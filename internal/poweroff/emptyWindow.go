package poweroff

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/poweroff/internal/monitorInfo"
)

type EmptyWindow struct {
	mainWindow *MainWindow
	gtk        *gtk.Window
}

// createBlankWindow creates a blank window to block extra screens
// other than the main window.
func createBlankWindow(mainWindow *MainWindow, monitor monitorInfo.MonitorInfo) (*EmptyWindow, error) {
	aw, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	win := &EmptyWindow{mainWindow: mainWindow, gtk: aw}

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

	win.gtk.Connect("key-press-event", win.keyPressed)

	return win, nil
}

// closeWindow  destroys the EmptyWindow
func (w *EmptyWindow) closeWindow() {
	w.gtk.Close()
	w.gtk.Destroy()
}

// keyPressed handles the key-press-event signal
func (w *EmptyWindow) keyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	w.mainWindow.keyPressed(nil, e)
}
