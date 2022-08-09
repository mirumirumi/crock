package render

import (
	"fmt"
	"strings"
	"time"

	print "github.com/fatih/color"

	"crock/color"
	"crock/terminal"
)

const (
	MIN_WIDTH = 120
	MIN_HEIGHT = 15
)

func Render(t terminal.Terminal, c color.Color, now time.Time) {
	char := char{now.Format("15:04:05")}
	crock := char.generate()

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
	numSpaceX := (t.Width - MIN_WIDTH) / 2 / 2 // 2nd `/2` is fine-tuning

	centered := []string{}
	for _, line := range lines {
		addedSpaces := strings.Repeat(" ", numSpaceX) + line
		centered = append(centered, addedSpaces)
	}
	xCenterd := strings.Join(centered, "\n")

	numSpaceY := (t.Height - MIN_HEIGHT) / 2 - 1
	yCenterd := strings.Repeat("\n", numSpaceY) + xCenterd + strings.Repeat("\n", numSpaceY)

	return yCenterd
}
