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
)

func Render(t terminal.Terminal, c color.Color, now time.Time) {
	crock := generate(now)

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

func generate(now time.Time) string {
	// formatted := now.Format("150405")
	// hour_1 := formatted[0:1]
	// hour_2 := formatted[1:2]
	// min_1 := formatted[2:3]
	// min_2 := formatted[3:4]
	// sec_1 := formatted[4:5]
	// sec_2 := formatted[5:6]






	return sample
}

func centering(crock string, t terminal.Terminal) string {
	lines := strings.Split(crock, "\n")
	numSpaces := (t.Width - MIN_WIDTH) / 2

	centered := []string{}
	for _, line := range lines {
		centered = append(centered, strings.Repeat(" ", numSpaces) + line + strings.Repeat(" ", numSpaces))
	}

	return strings.Join(centered, "\n")
}
