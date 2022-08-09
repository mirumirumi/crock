package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"crock/color"
	"crock/render"
	"crock/terminal"
)

const (
	MIN_TERM_WIDTH  = render.MIN_WIDTH
	MIN_TERM_HEIGHT = 15
)

var version = ""

func main() {
	var c color.Color = color.DEFAULT
	var err error

	if args := os.Args[1:]; 1 <= len(args) {
		c, err = parseArgs(args)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	var t terminal.Terminal
	t.New()

	if t.Width < MIN_TERM_WIDTH || t.Height < MIN_TERM_HEIGHT {
		needSize(t.Width, t.Height)
	}

	go crock(t, c)

	handleSignal(t)
}

func parseArgs(args []string) (color.Color, error) {
	switch {
	case args[0] == "--help" || args[0] == "-h":
		printHelp()
		os.Exit(0)
	case args[0] == "--version" || args[0] == "-v":
		printVersion()
		os.Exit(0)
	case args[0] == "--color" || args[0] == "-c":
		if 2 <= len(args) {
			c := args[1]
			switch c {
			case "black":
				return color.BLACK, nil
			case "white":
				return color.WHITE, nil
			case "green":
				return color.GREEN, nil
			case "yellow":
				return color.YELLOW, nil
			case "red":
				return color.RED, nil
			}
		} else {
			return color.DEFAULT, fmt.Errorf("failed to get the specified color, so it was set to default")
		}
	}

	return color.DEFAULT, nil
}

func needSize(x int, y int) {
	fmt.Println("\x1b[33mWoops!\x1b[0m crock needs a minimum size of", MIN_TERM_WIDTH, "(columns) x", MIN_TERM_HEIGHT, "(lines), please retry again.")
	fmt.Println("\nyour terminal size is:")
	fmt.Println(" x:", x)
	fmt.Println(" y:", y)
	fmt.Println("")
	os.Exit(0)
}

func printHelp() {
	fmt.Println("crock - which is rock clock 🪨")
}

func printVersion() {
	// go build xxx/xxx/ -ldflags "-s -w -X main.version=x.x.x"
	fmt.Println("crock: version ", getVersion())
}

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()

	if !ok {
		return "failed getting build version, please contact @mirumirumi at https://github.com/mirumirumi/crock/issues/new?title=Failed%20to%20get%20build%20version"
	}

	return info.Main.Version
}
