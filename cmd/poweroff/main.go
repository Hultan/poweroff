package main

import (
	"os"

	"github.com/hultan/poweroff/internal/poweroff"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	ApplicationId    = "se.softteam.poweroff"
	ApplicationFlags = glib.APPLICATION_FLAGS_NONE
)

func main() {
	// Create a new application
	application, err := gtk.ApplicationNew(ApplicationId, ApplicationFlags)
	if err != nil {
		panic("Failed to create GTK Application : " + err.Error())
	}

	powerOff := poweroff.PowerOff{}
	application.Connect("activate", powerOff.StartApplication)

	// Start the application (and exit when it is done)
	os.Exit(application.Run(nil))
}
