package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"crock/color"
	r "crock/render"
	"crock/terminal"
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

	if t.Width < r.MIN_WIDTH || t.Height < r.MIN_HEIGHT {
		needSize(t.Width, t.Height)
		os.Exit(0)
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
	msg := fmt.Sprintf(`
%sWoops!%s crock needs a minimum size of %d(columns) x %d(lines), please retry again.

your terminal size is:
 x: %d
 y: %d
`, "\x1b[33m", "\x1b[0m", r.MIN_WIDTH, r.MIN_HEIGHT, x, y)

	fmt.Println(msg)
}

func printHelp() {
	msg := fmt.Sprintf(`
crock - which is rock clock 🪨

Usage: crock [options] [args]

Options:
  --help   , -h    Show this help message
  --version, -v    Print the current crock version
  --color  , -c    Specify the output color for crock, you must specify the color after '--color' or '-c'
                     Valid colors:
                       black
                       white
                       green
                       yellow
                       red
`)

	fmt.Println(msg)
}

func printVersion() {
	// go build -ldflags "-s -w -X main.version=x.x.x"
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
