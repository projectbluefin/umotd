package internal

import "github.com/leonelquinteros/gotext"

func Usage(l *gotext.Locale) {
	println(l.Get("Usage:"))
	println("  umotd")
	println("  umotd config-path")
	println("  umotd tags (add <tag>... | remove <tag>... | list) ")
	println("  umotd version")
}
