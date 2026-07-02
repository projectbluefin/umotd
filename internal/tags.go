package internal

import (
	"github.com/leonelquinteros/gotext"
)

func TagsCommands(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No command specified."))
		return
	}
	switch args[0] {
	case "list":
		PrintTags()
	case "add":
		AddTags(args[1:], l)
	case "remove":
		RemoveTags(args[1:], l)
	default:
		println(l.Get("Command not recognized. Try 'list', 'add', or 'remove'."))
		return
	}
}

func PrintTags() {
	tags := ListTags()
	for _, tag := range tags {
		println(tag)
	}
}

func AddTags(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No tags to add."))
		return
	}
	for _, arg := range args {
		if err := AddTag(arg); err != nil {
			println(l.Get("Failed to add tag: %s", arg))
			println(l.Get("Error ~> %s", err.Error()))
		}
	}
}

func RemoveTags(args []string, l *gotext.Locale) {
	if len(args) < 1 {
		println(l.Get("No tags to remove."))
		return
	}
	for _, arg := range args {
		if err := RemoveTag(arg); err != nil {
			println(l.Get("Failed to remove tag: %s", arg))
			println(l.Get("Error ~> %s", err.Error()))
		}
	}
}
