package main

import (
	"os"
	"os/signal"

	"atomicgo.dev/cursor"

	"crock/terminal"
)

func handleSignal(t terminal.Terminal) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	cursor.Move(-(t.Width + LEN_STRFTIME), 0)
	cursor.Down(1)
	cursor.Show()
	os.Exit(0)
}
