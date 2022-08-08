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
	LEN_STRFTIME = 19
)

func crock(t terminal.Terminal, c color.Color) {
	tm.Clear()
	cursor.Hide()

	previous := ""

	for {
		var now string = time.Now().Format("2006/01/02 15:04:05")

		if previous == now {
			continue
		}

		tm.MoveCursor((t.Width - LEN_STRFTIME), 0)
		tm.Flush()

		fmt.Println(tm.Color(now, int(c)))
		render.Render(t, c)

		previous = now

		time.Sleep(INTERVAL)
	}
}
