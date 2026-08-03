package internal

import (
	"embed"
	"os"
	"strings"

	"golang.org/x/text/language"
)

// DetectLocale detects the current language
func DetectLocale(localesFS embed.FS) string {
	entries, err := localesFS.ReadDir("locales")
	if err != nil {
		return "en"
	}

	tags := []language.Tag{language.English}
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "en" {
			continue
		}
		if tag, err := language.Parse(entry.Name()); err == nil {
			tags = append(tags, tag)
		}
	}

	matcher := language.NewMatcher(tags)

	candidates := []string{}
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANGUAGE", "LANG"} {
		raw := os.Getenv(key)
		if raw == "" {
			continue
		}
		for _, part := range strings.Split(raw, ":") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			if idx := strings.Index(part, "."); idx >= 0 {
				part = part[:idx]
			}
			if idx := strings.Index(part, "@"); idx >= 0 {
				part = part[:idx]
			}
			candidates = append(candidates, strings.ReplaceAll(part, "_", "-"))
		}
	}

	for _, raw := range candidates {
		if tag, err := language.Parse(raw); err == nil {
			if match, _, _ := matcher.Match(tag); match != language.Und {
				return match.String()
			}
		}
	}

	return "en"
}
