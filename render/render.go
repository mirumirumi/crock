package render

import (
	"fmt"
	"strings"

	print "github.com/fatih/color"

	"crock/color"
	"crock/terminal"
)

const (
	CIRCLE_WIDTH = 52
)

func Render(t terminal.Terminal, c color.Color) {
	crock := circle







	crock = centering(crock, t)

	switch c {
	case color.BLACK:
		print.Black(crock)
	case color.WHITE:
		print.White(crock)
	case color.GREEN:
		print.Green(crock)
	case color.YELLOW:
		print.Yellow(crock)
	case color.RED:
		print.Red(crock)
	default:
		fmt.Println(crock)
	}
}

func centering(crock string, t terminal.Terminal) string {
	lines := strings.Split(crock, "\n")
	numSpaces := (t.Width - CIRCLE_WIDTH) / 2

	centered := []string{}
	for _, line := range lines {
		centered = append(centered, strings.Repeat(" ", numSpaces) + line + strings.Repeat(" ", numSpaces))
	}

	return strings.Join(centered, "\n")
}
