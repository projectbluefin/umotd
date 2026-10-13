package internal

import (
	"os"
	"os/exec"
	"regexp"
)

var OSName = getOSName()
var Desktop = getDesktop()

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

func getDesktop() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if desktop == "" {
		return "Unknown"
	}
	return desktop
}

func HasSystemApp(app string) bool {
	output, err := exec.LookPath(app)
	if err != nil {
		return false
	}
	return string(output) != ""
}
