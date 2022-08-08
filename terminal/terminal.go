package terminal

import (
	"runtime"

	tm "github.com/buger/goterm"
)

type OS int

const (
	LINUX OS = iota
	WINDOWS
	MAC
	OTHERS
)

type Terminal struct {
	Os     OS
	Width  int
	Height int
}

func (t *Terminal) New() {
	switch runtime.GOOS {
	case "linux":
		t.Os = LINUX
	case "windows":
		t.Os = WINDOWS
	case "darwin":
		t.Os = MAC
	default:
		t.Os = OTHERS
	}

	t.setSize()
}

func (t *Terminal) setSize() {
	t.Width = tm.Width()
	t.Height = tm.Height()
}
