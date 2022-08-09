package main

import (
	"fmt"
	"time"

	"atomicgo.dev/cursor"
	tm "github.com/buger/goterm"

	"crock/color"
	"crock/render"
	"crock/terminal"
)

const (
	INTERVAL     = 950 * time.Millisecond
	LEN_STRFTIME = 10
)

func crock(t terminal.Terminal, c color.Color) {
	tm.Clear()
	cursor.Hide()

	previous := ""

	for {
		var now = time.Now()
		var digital string = now.Format("01/02, 2006")

		if previous == digital {
			continue
		}

		tm.MoveCursor((t.Width - LEN_STRFTIME), 0)
		tm.Flush()

		fmt.Println(tm.Color(digital, int(c)))
		render.Render(t, c, now)

		previous = digital

		time.Sleep(INTERVAL)
	}
}
