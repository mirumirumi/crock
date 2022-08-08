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
	CIRCLE_WIDTH = 52
)

func Render(t terminal.Terminal, c color.Color, now time.Time) {
	crock := circle




	lines := strings.Split(crock, "\n")

	for _, line := range lines {
		for i, _ := range line {
			if line[i] != (" ") {
				;
			} else if h_0[i] != " " {
				
			} else if m_25[i] != " " {

			}


		}
	}








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
		centered = append(centered, strings.Repeat(" ", numSpaces - 2) + line + strings.Repeat(" ", numSpaces))  // 2: fine-tuning for the senses
	}

	return strings.Join(centered, "\n")
}
