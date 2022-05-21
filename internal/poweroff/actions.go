package poweroff

import (
	"fmt"
	"os/exec"

	"github.com/gotk3/gotk3/gtk"
)

type action struct {
	index    int
	name     string
	iconName string
	tooltip  string
	action   func()
}

func (m *MainForm) getActions() []action {
	return []action{
		{
			index:    0,
			name:     "<b>C</b>ancel",
			iconName: "cancel.png",
			tooltip:  "Exit back to Linux Mint...",
			action:   m.cancel,
		},
		{
			index:    1,
			name:     "<b>L</b>ock",
			iconName: "lock.png",
			tooltip:  "Lock the computer...",
			action:   m.lock,
		},
		{
			index:    2,
			name:     "L<b>o</b>gout",
			iconName: "logout.png",
			tooltip:  "Logout the current user session...",
			action:   m.logout,
		},
		{
			index:    3,
			name:     "<b>S</b>uspend",
			iconName: "suspend.png",
			tooltip:  "Suspend the computer...",
			action:   m.suspend,
		},
		{
			index:    4,
			name:     "<b>H</b>ibernate",
			iconName: "hibernate.png",
			tooltip:  "Hibernate the computer...",
			action:   m.hibernate,
		},
		{
			index:    5,
			name:     "<b>P</b>ower off",
			iconName: "powerOff.png",
			tooltip:  "Turn the computer off...",
			action:   m.powerOff,
		},
		{
			index:    6,
			name:     "<b>R</b>eboot",
			iconName: "reboot.png",
			tooltip:  "Reboot the computer...",
			action:   m.reboot,
		},
	}
}

func (m *MainForm) cancel() {
	m.Window.Close()
	m.Window.Destroy()
	gtk.MainQuit()
}

func (m *MainForm) lock() {
	runCommand("cinnamon-screensaver-command", "--lock")
	m.cancel()
}

func (m *MainForm) logout() {
	runCommand("cinnamon-session-quit", "--logout", "--force")
	m.cancel()
}

func (m *MainForm) suspend() {
	runCommand("systemctl", "suspend")
	m.cancel()
}

func (m *MainForm) hibernate() {
	err := runCommand("systemctl", "hibernate")
	if err != nil {
		m.suspend()
	}
	m.cancel()
}

func (m *MainForm) powerOff() {
	runCommand("systemctl", "poweroff")
	m.cancel()
}

func (m *MainForm) reboot() {
	runCommand("systemctl", "reboot")
	m.cancel()
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
