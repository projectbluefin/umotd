package internal

import (
	"os"
	"os/exec"
	"regexp"
)

var OSName = getOSName()
var ImgName = getImageName()
var Desktop = getDesktop()

// getImageName gets the Image name from image-info.json
func getImageName() string {
	data, err := os.ReadFile("/usr/share/ublue-os/image-info.json")
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`"image-name": "(.*)"`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// getOSName gets the OS name from os-release
func getOSName() string {
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

// getDesktop gets the desktop environment name the environment variable $XDG_CURRENT_DESKTOP
func getDesktop() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if desktop == "" {
		return "Unknown"
	}
	return desktop
}

/*
HasSystemApp returns if exec.LookPath returns the path of the targeted app or not.

It works with embedded commands or homebrew apps, but not Flatpaks.
*/
func HasSystemApp(app string) bool {
	output, err := exec.LookPath(app)
	if err != nil {
		return false
	}
	return string(output) != ""
}
