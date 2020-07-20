// +build !windows,!android,!ios

package dialog

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (f *fileDialog) loadPlaces() []fyne.CanvasObject {
	return []fyne.CanvasObject{widget.NewButtonWithIcon("Computer", theme.ComputerIcon(), func() {
		f.setDirectory("/")
	})}
}

func isHidden(file, _ string) bool {
	return len(file) == 0 || file[0] == '.'
}

func fileOpenOSOverride(*FileDialog) bool {
	return false
}

func fileSaveOSOverride(*FileDialog) bool {
	return false
}
