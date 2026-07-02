package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Tags []string `json:"tags"`
}

func AddTag(newTag string) error {
	cfg := GetConfig()
	cfg.Tags = append(cfg.Tags, newTag)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetPath(), data, 0644)
}

func RemoveTag(tagToRemove string) error {
	cfg := GetConfig()
	for i, tag := range cfg.Tags {
		if tag == tagToRemove {
			cfg.Tags = append(cfg.Tags[:i], cfg.Tags[i+1:]...)
			break
		}
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(GetPath(), data, 0644)
}

func ListTags() []string {
	if cfg := GetConfig(); cfg.Tags != nil {
		return cfg.Tags
	}
	return nil
}

func GetPath() string {
	return "/etc/ublue-os/tags.json"
}

// GetConfig returns the config file at the given path, or a default config if no valid config file is found
func GetConfig() Config {
	cfg := Config{}
	path := GetPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg
	}

	return cfg
}
