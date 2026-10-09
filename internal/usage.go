package internal

import (
	"fmt"

	"github.com/leonelquinteros/gotext"
)

const space = "          "

// Usage
func Usage(l *gotext.Locale) {
	fmt.Println(l.Get("uMotd is a `Messages Of The Day` system for Universal Blue operating systems.") + "\n")
	fmt.Println(l.Get("Usage:") + "\n")
	fmt.Println(space + "umotd")
	fmt.Println(space + "umotd tags (add <tag>... | remove <tag>... | list) ")
	fmt.Println(space + "umotd config-path")
	fmt.Println(space + "umotd version")
}
