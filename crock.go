package main

import (
	"fmt"
	"strings"
	"time"

	"atomicgo.dev/cursor"
	tm "github.com/buger/goterm"

	"crock/color"
	"crock/render"
	"crock/terminal"
)

const (
	INTERVAL     = 950 * time.Millisecond
	LEN_STRFDATE = 11
)

func crock(t terminal.Terminal, c color.Color) {
	tm.Clear()
	cursor.Hide()

	previous := ""

	for {
		var now = time.Now()
		var digital string = now.Format("01/02, 2006")
		var sec string = now.Format("2006/01/02 15:04:05")

		if previous == sec {
			continue
		}

		title := "crock🪨 is crocking:"
		first := title + strings.Repeat(" ", t.Width - len(title) - LEN_STRFDATE) + digital

		tm.MoveCursor(0, 0)
		tm.Flush()

		fmt.Println(tm.Color(first, int(c)))
		render.Render(t, c, now)

		previous = sec

		time.Sleep(INTERVAL)
	}
}
