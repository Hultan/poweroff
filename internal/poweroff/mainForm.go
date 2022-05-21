package poweroff

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

const applicationTitle = "poweroff"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	AboutDialog *gtk.AboutDialog
}

var supportsAlpha bool

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Get the main window from the glade file
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m.Window = win

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("poweroff main window")
	m.Window.SetDecorated(false)
	m.Window.Maximize()
	m.Window.SetName("mainWindow")

	// Hook up events
	m.Window.Connect("destroy", m.Cancel)
	m.Window.Connect("key-press-event", m.KeyPressed)
	// m.Window.Connect("draw", m.Draw)
	// m.Window.Connect("screen-changed", m.ScreenChanged)

	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}
	grid.SetMarginStart(10)
	grid.SetMarginTop(10)
	grid.SetColumnSpacing(10)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SetVAlign(gtk.ALIGN_CENTER)

	m.AddButton(grid, 0, "Cancel", "/home/per/code/poweroff/assets/cancel.png", m.Cancel)
	m.AddButton(grid, 1, "Lock", "/home/per/code/poweroff/assets/lock.png", m.Cancel)
	m.AddButton(grid, 2, "Logout", "/home/per/code/poweroff/assets/logout.png", m.Cancel)
	m.AddButton(grid, 3, "Suspend", "/home/per/code/poweroff/assets/suspend.png", m.Cancel)
	m.AddButton(grid, 4, "Hibernate", "/home/per/code/poweroff/assets/hibernate.png", m.Cancel)
	m.AddButton(grid, 5, "Shutdown", "/home/per/code/poweroff/assets/shutdown.png", m.Cancel)
	m.AddButton(grid, 6, "Restart", "/home/per/code/poweroff/assets/restart.png", m.Cancel)

	TransparentBackground(win)
	m.Window.Add(grid)

	// Show the main window
	m.Window.ShowAll()
	gtk.Main()
}

func TransparentBackground(win *gtk.ApplicationWindow) (err error) {
	var (
		css = `
#mainWindow {
	background-color: rgba(255, 255, 255, 0.2);
}
label {
	color: #AAAAAA;
	font-size: 24px;
}
`
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
	screen = win.GetScreen()
	if visual, err = screen.GetRGBAVisual(); err == nil {
		win.SetVisual(visual)
	}
	return
}

func (m *MainForm) Cancel() {
	m.Window.Close()
	m.Window.Destroy()
	gtk.MainQuit()
}

func (m *MainForm) KeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)
	fmt.Println("Key pressed:", ke.KeyVal())

	switch ke.KeyVal() {
	case gdk.KEY_Q:
		m.Cancel()
	case gdk.KEY_q:
		m.Cancel()
	case gdk.KEY_Escape:
		m.Cancel()
	}
}

func (m *MainForm) AddButton(grid *gtk.Grid, index int, text string, path string, action func()) error {
	image, err := gtk.ImageNewFromFile(path)
	if err != nil {
		return err
	}
	btn, err := gtk.ButtonNew()
	if err != nil {
		return err
	}
	btn.SetImage(image)
	btn.Connect("clicked", action)

	lbl, err := gtk.LabelNew(text)
	if err != nil {
		return err
	}
	lbl.SetName("buttonLabel")
	lbl.SetMarkup("<b>" + text + "</b>")

	grid.Attach(btn, index, 0, 1, 1)
	grid.Attach(lbl, index, 1, 1, 1)

	return nil
}
