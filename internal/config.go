package internal

import (
	"encoding/json"
	"os"
)

const TagsPath = "/etc/ublue-os/tags.json"

var Cfg = getConfig()

type Config struct {
	Tags []string `json:"tags"`
}

// AddTag adds the selected tag(s) to the config
func AddTag(newTag string) error {
	Cfg.Tags = append(Cfg.Tags, newTag)
	data, err := json.MarshalIndent(Cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(TagsPath, data, 0644)
}

// RemoveTag removes the selected tag(s) from the config
func RemoveTag(tagToRemove string) error {
	for i, tag := range Cfg.Tags {
		if tag == tagToRemove {
			Cfg.Tags = append(Cfg.Tags[:i], Cfg.Tags[i+1:]...)
			break
		}
	}
	data, err := json.MarshalIndent(Cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(TagsPath, data, 0644)
}

// ListTags lists all the tags of the config
func ListTags() []string {
	if Cfg.Tags != nil {
		return Cfg.Tags
	}
	return nil
}

// getConfig returns the current configuration
func getConfig() Config {
	config := Config{}

	data, err := os.ReadFile(TagsPath)
	if err != nil {
		return config
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return config
	}

	return config
}
