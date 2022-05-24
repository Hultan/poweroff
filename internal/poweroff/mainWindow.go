package poweroff

import (
	"path"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/poweroff/internal/monitorInfo"
)

type mainWindow struct {
	app          *gtk.Application
	gtk          *gtk.ApplicationWindow
	grid         *gtk.Grid
	monitors     []monitorInfo.MonitorInfo
	emptyWindows []*blankWindow
}

func createMainWindow(app *gtk.Application, monitors []monitorInfo.MonitorInfo) (*mainWindow, error) {
	aw, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		return nil, err
	}
	win := &mainWindow{app: app, gtk: aw}

	// Set up main window
	win.gtk.SetDecorated(false)
	win.gtk.SetName("window")
	win.gtk.Connect("destroy", gtk.MainQuit)
	win.gtk.Connect("key-press-event", win.keyPressed)

	err = makeWindowTransparent(win.gtk)
	if err != nil {
		return nil, err
	}

	err = win.addActionButtons()
	if err != nil {
		return nil, err
	}

	// Show the main window
	win.gtk.ShowAll()
	win.gtk.Fullscreen()
	win.gtk.Move(monitors[0].Left, monitors[0].Top)
	win.gtk.SetKeepAbove(true)
	// win.gtk.SetCanFocus(true)
	// win.gtk.GrabDefault()

	// Add blank blocking windows for monitors other than the main monitor
	for i := 1; i < len(monitors); i++ {
		ew, err := createBlankWindow(monitors[i])
		if err != nil {
			return nil, err
		}
		win.emptyWindows = append(win.emptyWindows, ew)
	}

	return win, nil
}

// closeApplication destroys the emptyWindows and destroys the mainWindow
func (w *mainWindow) closeApplication() {
	for _, e := range w.emptyWindows {
		e.closeWindow()
	}

	w.gtk.Close()
	w.gtk.Destroy()
}

// addActionButtons adds a button for each action
func (w *mainWindow) addActionButtons() error {
	grid, err := gtk.GridNew()
	if err != nil {
		return err
	}

	grid.SetName("grid")
	grid.SetColumnSpacing(20)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SetVAlign(gtk.ALIGN_CENTER)

	actions := w.getActions()
	for _, a := range actions {
		err = w.addButton(grid, a)
		if err != nil {
			return err
		}
	}

	w.grid = grid
	w.gtk.Add(grid)

	return nil
}

// addButton adds a new action button
func (w *mainWindow) addButton(grid *gtk.Grid, a action) error {
	p := path.Join("/home/per/code/poweroff/assets", a.iconName)
	image, err := gtk.ImageNewFromFile(p)
	if err != nil {
		return err
	}

	btn, err := gtk.ButtonNew()
	if err != nil {
		return err
	}
	btn.SetImage(image)
	btn.Connect("clicked", a.action)
	btn.SetTooltipText(a.tooltip)

	lbl, err := gtk.LabelNew(a.name)
	if err != nil {
		return err
	}
	lbl.SetName("buttonLabel")
	lbl.SetMarkup(a.name)

	grid.Attach(btn, a.index, 0, 1, 1)
	grid.Attach(lbl, a.index, 1, 1, 1)

	return nil
}

// keyPressed handles the key-press-event signal
func (w *mainWindow) keyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	switch ke.KeyVal() {
	case gdk.KEY_C, gdk.KEY_c, gdk.KEY_Q, gdk.KEY_q, gdk.KEY_Escape:
		w.gtk.Destroy()
	case gdk.KEY_L, gdk.KEY_l:
		w.lock()
	case gdk.KEY_O, gdk.KEY_o:
		w.logout()
	case gdk.KEY_S, gdk.KEY_s:
		w.suspend()
	case gdk.KEY_H, gdk.KEY_h:
		w.hibernate()
	case gdk.KEY_P, gdk.KEY_p:
		w.powerOff()
	case gdk.KEY_R, gdk.KEY_r:
		w.reboot()
	}
}
