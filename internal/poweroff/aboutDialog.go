package poweroff

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

func (m *MainForm) openAboutDialog(fw *framework.Framework) {
	if m.AboutDialog == nil {
		about := m.builder.GetObject("about_dialog").(*gtk.AboutDialog)
		about.SetDestroyWithParent(true)
		about.SetTransientFor(m.Window)
		about.SetProgramName(applicationTitle)
		about.SetComments("A Linux poweroff application written in GTK and Go.")
		about.SetVersion(applicationVersion)
		about.SetCopyright(applicationCopyRight)
		image, err := gdk.PixbufNewFromFile(fw.Resource.GetResourcePath("application.png"))
		if err == nil {
			about.SetLogo(image)
		}
		about.SetModal(true)
		about.SetPosition(gtk.WIN_POS_CENTER)

		about.Connect("response", func(dialog *gtk.AboutDialog, responseId gtk.ResponseType) {
			if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
				about.Hide()
			}
		})

		m.AboutDialog = about
	}

	m.AboutDialog.Present()
}
