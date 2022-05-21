package poweroff

import (
	"fmt"
	"os"
	"path"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/poweroff/internal/softmonitorInfo"
	"github.com/hultan/softteam/framework"
)

const applicationTitle = "poweroff"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	AboutDialog *gtk.AboutDialog
	extra       []*gtk.ApplicationWindow
}

var supportsAlpha bool

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) (err error) {
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
	m.Window.SetDecorated(false)
	m.Window.SetName("mainWindow")

	// Hook up events
	m.Window.Connect("destroy", m.cancel)
	m.Window.Connect("key-press-event", m.KeyPressed)

	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}
	grid.SetName("grid")
	// grid.SetBorderWidth(10)
	// grid.SetMarginStart(10)
	// grid.SetMarginEnd(10)
	// grid.SetMarginTop(10)
	// grid.SetMarginBottom(10)
	grid.SetColumnSpacing(20)
	grid.SetHAlign(gtk.ALIGN_CENTER)
	grid.SetVAlign(gtk.ALIGN_CENTER)

	actions := m.getActions()
	for _, a := range actions {
		m.AddButton(grid, a)
	}

	TransparentBackground(win)
	m.Window.Add(grid)

	info := softmonitorInfo.NewSoftMonitorInfo()
	monitors, err := info.GetMonitorInfo()

	// Dim extra monitors
	if len(monitors) > 1 {
		for i := 1; i < len(monitors); i++ {
			monitor := monitors[i]
			w := m.createWindow(app, monitor)
			m.extra = append(m.extra, w)
			fmt.Printf("%d : %dx%d + %dx%d\n", monitor.Number, monitor.Width, monitor.Height, monitor.Top, monitor.Left)
		}
	}

	// Show the main window
	m.Window.ShowAll()
	m.Window.Move(monitors[0].Left, monitors[0].Top)
	m.Window.Fullscreen()
	m.Window.SetKeepAbove(true)

	gtk.Main()

	return nil
}

func (m *MainForm) KeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	ke := gdk.EventKeyNewFromEvent(e)

	// fmt.Println("Key pressed:", ke.KeyVal())

	switch ke.KeyVal() {
	case gdk.KEY_C, gdk.KEY_c, gdk.KEY_Q, gdk.KEY_q, gdk.KEY_Escape:
		m.cancel()
	case gdk.KEY_L, gdk.KEY_l:
		m.lock()
	case gdk.KEY_O, gdk.KEY_o:
		m.logout()
	case gdk.KEY_S, gdk.KEY_s:
		m.suspend()
	case gdk.KEY_H, gdk.KEY_h:
		m.hibernate()
	case gdk.KEY_P, gdk.KEY_p:
		m.powerOff()
	case gdk.KEY_R, gdk.KEY_r:
		m.reboot()
	}
}

func (m *MainForm) AddButton(grid *gtk.Grid, a action) (err error) {
	p := path.Join("/home/per/code/poweroff/assets", a.iconName)
	image, err := gtk.ImageNewFromFile(p)
	if err != nil {
		return
	}

	btn, err := gtk.ButtonNew()
	if err != nil {
		return
	}
	btn.SetImage(image)
	btn.Connect("clicked", a.action)
	btn.SetTooltipText(a.tooltip)

	lbl, err := gtk.LabelNew(a.name)
	if err != nil {
		return
	}
	lbl.SetName("buttonLabel")
	lbl.SetMarkup(a.name)

	grid.Attach(btn, a.index, 0, 1, 1)
	grid.Attach(lbl, a.index, 1, 1, 1)

	return
}

func TransparentBackground(win *gtk.ApplicationWindow) (err error) {
	var (
		css = `
#mainWindow {
	background-color: rgba(255, 255, 255, 0.2);
}
#grid {
	background-color : #222222;
	border-top : 20px solid #222222;
	border-right : 20px solid #222222;
	border-bottom : 10px solid #222222;
	border-left : 20px solid #222222;
	border-top-right-radius: 25px;
	border-bottom-right-radius: 25px;
	border-top-left-radius: 25px;
	border-bottom-left-radius: 25px;
}
button {
	border : 2px solid #666666;
	background-color : #444444;
	border-top-right-radius: 10px;
	border-bottom-right-radius: 10px;
	border-top-left-radius: 10px;
	border-bottom-left-radius: 10px;
}
button:hover {
    background: #559955;
    border-radius: 100px;
    transition: all 1s ease;
}
#buttonLabel {
	background-color : #222222;
	margin-top : 10px;
	color: #AAAAAA;
	font-size: 20px;
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

func (m *MainForm) createWindow(app *gtk.Application, monitor softmonitorInfo.MonitorInfo) *gtk.ApplicationWindow {
	// Get the main window from the glade file
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Set up main window
	win.SetDecorated(false)
	win.SetName("mainWindow")
	win.SetSizeRequest(monitor.Width, monitor.Height)
	win.Move(monitor.Left, monitor.Top)
	win.Fullscreen()
	TransparentBackground(win)
	win.ShowAll()
	win.SetKeepAbove(true)
	return win
}
