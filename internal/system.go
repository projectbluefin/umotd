package internal

import (
	"os"
	"regexp"
)

var OSName = getOSName()

func getOSName() string {
	// Gets the OS name from /etc/os-release
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`NAME="(.*)"`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		return match[1]
	}
	return "Your System"
}
