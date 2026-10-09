package main

import (
	"fmt"
	"os"

	i "umotd/internal"

	"github.com/leonelquinteros/gotext"
)

const version = "0.3.3"

func main() {

	// Loads the locale based on the system's locale

	l := gotext.NewLocale("locales", i.DetectLocale())
	l.AddDomain("default")

	// Handles command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		// Returns the path of the current config file
		case "config-path":
			fmt.Println(i.TagsPath)
			return
		// Redirects to tag related commands
		case "tags":
			i.TagsCommands(os.Args[2:], l)
			return
		// Prints the version
		case "version":
			fmt.Println(version)
			return
		// Shows the usage
		case "help":
			i.Usage(l)
			return
		default:
			fmt.Fprintln(os.Stderr, l.Get("Invalid command."))
			i.Usage(l)
			return
		}
	}

	// Loads the configuration from the system's config file

	fmt.Print(i.GetRandomMessage(i.Cfg.Tags, l))
}
