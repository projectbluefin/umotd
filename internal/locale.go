package internal

import (
	"embed"
	"os"
	"slices"
	"strings"

	"golang.org/x/text/language"
)

// Most of the code here has been taken and tweaked from the `go-locale` package.

func getUnixLocales() []string {
	locale := getLangFromEnv()
	if locale == "C" || locale == "POSIX" || len(locale) == 0 {
		return nil
	}

	return splitLocales(locale)
}

func splitLocales(locales string) []string {
	// If the user set different locales, they might be set in $LANGUAGE,
	// separated by a colon
	return strings.Split(locales, ":")
}

func splitLocale(locale string) (string, string) {
	// Remove the encoding, if present.
	formattedLocale, _, _ := strings.Cut(locale, ".")

	// Normalize by replacing the hyphens with underscores
	formattedLocale = strings.ReplaceAll(formattedLocale, "-", "_")

	// Split at the underscore.
	language, territory, _ := strings.Cut(formattedLocale, "_")
	return language, territory
}

func getLangFromEnv() string {
	locale := ""

	// Check the following environment variables for the language information
	// See https://www.gnu.org/software/gettext/manual/html_node/Locale-Environment-Variables.html
	for _, env := range [...]string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		locale = os.Getenv(env)
		if len(locale) > 0 {
			break
		}
	}

	if locale == "C" || locale == "POSIX" {
		return locale
	}

	// Check LANGUAGE if localization is enabled (not set to "C")
	// See https://www.gnu.org/software/gettext/manual/html_node/The-LANGUAGE-variable.html#The-LANGUAGE-variable
	languages := os.Getenv("LANGUAGE")
	if len(languages) > 0 {
		return languages
	}

	return locale
}

func GetLocale(localesFS embed.FS) string {
	unixLocales := getUnixLocales()

	if unixLocales == nil {
		return ""
	}

	entries, err := localesFS.ReadDir("locales")
	if err != nil {
		return ""
	}

	tags := []language.Tag{language.English}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if tag, err := language.Parse(entry.Name()); err == nil {
			tags = append(tags, tag)
		}
	}

	for _, raw := range unixLocales {
		locale, region := splitLocale(raw)

		if region != "" {
			if localeParsed, err := language.Parse(locale + "-" + region); err == nil && slices.Contains(tags, localeParsed) {
				return locale + "_" + region
			}
		}

		if localeParsed, err := language.Parse(locale); err == nil && slices.Contains(tags, localeParsed) {
			return locale
		}
	}

	return "en"
}
