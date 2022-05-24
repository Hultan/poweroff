package poweroff

import (
	"fmt"
	"os/exec"
)

type action struct {
	index    int
	name     string
	iconName string
	tooltip  string
	action   func()
}

func (w *mainWindow) getActions() []action {
	return []action{
		{
			index:    0,
			name:     "<b>C</b>ancel",
			iconName: "cancel.png",
			tooltip:  "Exit back to Linux Mint...",
			action:   w.gtk.Destroy,
		},
		{
			index:    1,
			name:     "<b>L</b>ock",
			iconName: "lock.png",
			tooltip:  "Lock the computer...",
			action:   w.lock,
		},
		{
			index:    2,
			name:     "L<b>o</b>gout",
			iconName: "logout.png",
			tooltip:  "Logout the current user session...",
			action:   w.logout,
		},
		{
			index:    3,
			name:     "<b>S</b>uspend",
			iconName: "suspend.png",
			tooltip:  "Suspend the computer...",
			action:   w.suspend,
		},
		{
			index:    4,
			name:     "<b>H</b>ibernate",
			iconName: "hibernate.png",
			tooltip:  "Hibernate the computer...",
			action:   w.hibernate,
		},
		{
			index:    5,
			name:     "<b>P</b>ower off",
			iconName: "powerOff.png",
			tooltip:  "Turn the computer off...",
			action:   w.powerOff,
		},
		{
			index:    6,
			name:     "<b>R</b>eboot",
			iconName: "reboot.png",
			tooltip:  "Reboot the computer...",
			action:   w.reboot,
		},
	}
}

func (w *mainWindow) lock() {
	err := w.runCommand("cinnamon-screensaver-command", "--lock")
	if err != nil {
		fmt.Println(err)
	}
	w.gtk.Destroy()
}

func (w *mainWindow) logout() {
	err := w.runCommand("cinnamon-session-quit", "--logout", "--force")
	if err != nil {
		fmt.Println(err)
	}
	w.gtk.Destroy()
}

func (w *mainWindow) suspend() {
	err := w.runCommand("systemctl", "suspend")
	if err != nil {
		fmt.Println(err)
	}
	w.gtk.Destroy()
}

func (w *mainWindow) hibernate() {
	err := w.runCommand("systemctl", "hibernate")
	if err != nil {
		w.suspend()
	}
	w.gtk.Destroy()
}

func (w *mainWindow) powerOff() {
	err := w.runCommand("systemctl", "poweroff")
	if err != nil {
		fmt.Println(err)
	}
	w.gtk.Destroy()
}

func (w *mainWindow) reboot() {
	err := w.runCommand("systemctl", "reboot")
	if err != nil {
		fmt.Println(err)
	}
	w.gtk.Destroy()
}

func (w *mainWindow) runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
