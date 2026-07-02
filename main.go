package main

import (
	"embed"
	"fmt"
	"os"

	"umotd/internal"

	"github.com/leonelquinteros/gotext"
)

const VERSION = "0.3"

//go:embed all:locales
var localesFS embed.FS

func main() {

	// Loads the locale based on the system's locale
	locale := internal.DetectLocale(localesFS)
	l := gotext.NewLocaleFSWithPath(locale, localesFS, "locales")
	l.AddDomain("default")

	// Handles command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {

		// Prints the version
		case "--version", "-v", "version":
			fmt.Println(VERSION)
			return
		// Returns the path to the current file
		case "config-path":
			fmt.Println(internal.GetPath())
			return
		case "tags":
			internal.TagsCommands(os.Args[2:], l)
			return
		default:
			fmt.Println(l.Get("Invalid command"))
			return
		}
	}

	// Loads the configuration from the system's config file
	cfg := internal.GetConfig()

	fmt.Print(internal.GetRandomMessage(l, cfg.Tags))
}
