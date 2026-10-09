package internal

import (
	"fmt"

	"github.com/leonelquinteros/gotext"
)

const space = "          "

// Usage
func Usage(l *gotext.Locale) {
	fmt.Println(l.Get("uMOTD is a translatable set of Messages Of The Day for Universal Blue systems.") + "\n")
	fmt.Println(l.Get("Usage:") + "\n")
	fmt.Println(space + "umotd")
	fmt.Println(space + "umotd tags (add <tag>... | remove <tag>... | list) ")
	fmt.Println(space + "umotd config-path")
	fmt.Println(space + "umotd version")
}
