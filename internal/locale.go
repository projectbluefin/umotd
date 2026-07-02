package internal

import (
	"embed"
	"os"
	"strings"

	"golang.org/x/text/language"
)

// Language detection
func DetectLocale(localesFS embed.FS) string {
	langDir, _ := localesFS.ReadDir("locales")

	tags := []language.Tag{language.English}
	for _, languageDir := range langDir {
		if languageDir.Name() == "en" {
			continue
		}
		tags = append(tags, language.Make(languageDir.Name()))
	}

	var supported = language.NewMatcher(tags)

	raw := os.Getenv("LANGUAGE")
	if raw == "" {
		raw = os.Getenv("LANG")
	}
	if raw == "" {
		raw = os.Getenv("LC_ALL")
	}
	raw = strings.Split(raw, ".")[0]
	raw = strings.Replace(raw, "_", "-", 1)

	tag := language.Make(raw)
	match, _, _ := supported.Match(tag)

	base, _ := match.Base()
	return base.String()
}
