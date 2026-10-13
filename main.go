package main

import (
	"fmt"
	"os"

	i "umotd/internal"

	"github.com/leonelquinteros/gotext"
)

const version = "0.4.0"

func main() {

	// Loads the locale based on the system's locale

	l := gotext.NewLocale("locales", i.DetectLocale())
	l.AddDomain("default")

	// Handles command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		// Prints the version
		case "version":
			fmt.Println(version)
			return
		// Shows the usage
		case "help":
			i.Usage(l)
			return
		// Prints a special message for Fedora-based isos
		case "fb-iso":
			fmt.Print(l.Get("󰋼 The Live ISO is designed for installation and troubleshooting. `rpm-ostree` commands will not work here. If you've successfully installed %s, please restart your computer.", i.OSName))
			return
		default:
			fmt.Fprintln(os.Stderr, l.Get("Invalid command."))
			i.Usage(l)
			return
		}
	}

	// Prints a random message that's relevant to the system

	fmt.Print(i.GetRandomMessage(l))
}
